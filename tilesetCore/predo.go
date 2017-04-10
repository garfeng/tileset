package main

import (
	"image"
	"image/color"
)

type PreDo struct {
	src image.Image
	Hue bool
}

func newPreDo(src image.Image, hue bool) *PreDo {
	return &PreDo{src, hue}
}

func newPreDoAndSave(src image.Image, hue bool, output string) error {
	p := newPreDo(src, hue)
	return savePng(p, output)
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
