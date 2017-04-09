package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func parseCenter(a, b, x int) int {
	if x < a {
		x = a
	} else if x > b {
		x = b
	}
	return x
}

type PreDo struct {
	src image.Image
}

func (pd *PreDo) At(x, y int) color.Color {
	var (
	//DEF = 65
	//PD  = 6
	)
	x1, y1 := x%64, y%64
	minX, minY := (x/64)*32, (y/64)*32
	maxX, maxY := minX+31, minY+31
	x2, y2 := x1+minX-16, y1+minY-16
	x3, y3 := parseCenter(minX, maxX, x2), parseCenter(minY, maxY, y2)
	//fmt.Println(x3, y3)
	return pd.src.At(x3, y3)
}

func (pd *PreDo) Bounds() image.Rectangle {
	bd := pd.src.Bounds()
	return image.Rect(0, 0, bd.Dx()*2, bd.Dy()*2)
}

func (pd *PreDo) ColorModel() color.Model {
	return pd.src.ColorModel()
}

type AftDo struct {
	src image.Image
}

func (ad *AftDo) At(x, y int) color.Color {
	x1, y1 := x%48, y%48
	x2, y2 := x/48, y/48
	x3, y3 := x2*96+24+x1, y2*96+24+y1
	fmt.Println(x, y, "=>", x3, y3)
	return ad.src.At(x3, y3)
}

func (ad *AftDo) Bounds() image.Rectangle {
	bd := ad.src.Bounds()
	return image.Rect(0, 0, bd.Dx()/2, bd.Dy()/2)
}

func (ad *AftDo) ColorModel() color.Model {
	return ad.src.ColorModel()
}

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
