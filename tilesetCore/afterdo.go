package main

import (
	"image"
	"image/color"
)

type AftDo struct {
	src image.Image
}

func newAftDo(img image.Image) *AftDo {
	return &AftDo{img}
}

func newAftDoAndSave(img image.Image, output string) error {
	p := newAftDo(img)
	return savePng(p, output)
}

func (ad *AftDo) At(x, y int) color.Color {
	x1, y1 := x%48, y%48
	x2, y2 := x/48, y/48
	x3, y3 := x2*96+24+x1, y2*96+24+y1
	return ad.src.At(x3, y3)
}

func (ad *AftDo) Bounds() image.Rectangle {
	bd := ad.src.Bounds()
	return image.Rect(0, 0, bd.Dx()/2, bd.Dy()/2)
}

func (ad *AftDo) ColorModel() color.Model {
	return ad.src.ColorModel()
}
