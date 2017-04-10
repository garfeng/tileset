package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"code.aliyun.com/JRY/mtquery/module/mfile"
)

func savePng(src image.Image, filename string) error {
	fp, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fp.Close()
	return png.Encode(fp, src)
}

func readPng(filename string) (image.Image, error) {
	w, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer w.Close()
	return png.Decode(w)
}

func isPng(filename string) bool {
	pngExp, err := regexp.Compile(".png$")
	if err != nil {
		return false
	}
	return pngExp.MatchString(filename)
}

func parseCenter(a, b, x int) int {
	if x < a {
		x = a
	} else if x > b {
		x = b
	}
	return x
}

func tilesetCore(in, out, c string, hue bool) {
	if mfile.IsFile(in) {
		if mfile.IsDir(out) {
			_, inName := filepath.Split(in)
			out = filepath.Join(out, inName)
		}
		out = deleteExt(out)
		err := handleSingleImg(in, out, c, hue)
		if err != nil {
			fmt.Println(err)
		}
	} else if mfile.IsDir(in) {
		if mfile.IsFile(out) {
			fmt.Println("If input is dir, output must be dir also.")
			return
		} else if !mfile.Exist(out) {
			os.MkdirAll(out, 0755)
		}
		handleDirImg(in, out, c, hue)
	}
}

func handleSingleImg(in, out, c string, hue bool) error {
	fmt.Println("\n\nHandle", in, "...")
	img, err := readPng(in)
	if err != nil {
		return err
	}
	needSplit := isXp(img)
	if needSplit {
		imageList := newxpTile(img)
		fmt.Println("split to", len(imageList), "images")
		for i, v := range imageList {
			of := fmt.Sprintf("%s_%d.png", out, i+1)
			err := handleSingleSrc(v, of, c, hue)
			if err != nil {
				return err
			}
		}
	} else {
		of := fmt.Sprintf("%s.png", out)
		handleSingleSrc(img, of, c, hue)
	}
	return nil
}

const (
	TMP1 = "tmp1.png"
	TMP2 = "tmp2.png"
)

func handleSingleSrc(src image.Image, out, c string, hue bool) error {
	fmt.Println("Do something before resize")
	err := newPreDoAndSave(src, hue, TMP1)
	if err != nil {
		return err
	}
	fmt.Println("Resizing...")
	err = runWaifu2x(TMP1, TMP2, c)
	if err != nil {
		return err
	}
	img2, err := readPng(TMP2)
	if err != nil {
		return err
	}
	fmt.Println("Save to", out)
	return newAftDoAndSave(img2, out)
}

func runWaifu2x(in, out string, c string) error {
	if c != "gpu" {
		c = "cpu"
	}
	cmd := exec.Command("waifu2x-caffe-cui.exe", "-i", in, "-o", out, "-p", c, "-s", "1.5")
	return cmd.Run()
}

//handleDirImg(in, out, c, hue)

func handleDirImg(in, out, c string, hue bool) {
	list := mfile.ScanDir(in)
	for _, file := range list {
		if isPng(file) {
			inFile := filepath.Join(in, file)
			outFile := filepath.Join(out, file)
			outFile = deleteExt(outFile)
			err := handleSingleImg(inFile, outFile, c, hue)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
