package main

import (
	"os"
	"regexp"
)

/*
func main() {
	w, err := os.Open(os.Args[2])
	img, err := png.Decode(w)
	newFileName := "test"
	if err != nil {
		log.Fatal(err)
	}
	var (
		newImg image.Image
	)
	if os.Args[1] == "-p" {
		newImg = &PreDo{img}
		newFileName += "_p.png"
	} else if os.Args[1] == "-a" {
		newImg = &AftDo{img}
		newFileName += "_a.png"
	}
	w.Close()
	w2, err := os.Create(newFileName)
	png.Encode(w2, newImg)
	w2.Close()
}
*/

func main() {
	if len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		printHelpInDetail()
		return
	}
	args, err := parseArgs()
	if err != nil {
		printHelp()
		return
	}
	in := args.Get(INPUT)
	out := args.Get(OUTPUT)
	c := args.Get(CORE)
	hue := false
	xp_resize := false
	if args.Get(MODIFY_HUI) == "true" {
		hue = true
	}
	if args.Get(XP_RESIZE) == "true" {
		xp_resize = true
	}
	tilesetCore(in, out, c, hue, xp_resize)
}

func deleteExt(path string) string {
	extREG, _ := regexp.Compile(".png$")
	return extREG.ReplaceAllString(path, "")
}
