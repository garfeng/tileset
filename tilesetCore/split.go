package main

import (
	"image"
	"image/color"
)

/**
 * Split xp tilesets (256 x height) => (512 x 512)
 */

func isXp(img image.Image) bool {
	return img.Bounds().Dx() == 256 && img.Bounds().Dy() > 256
}

type xpTile struct {
	src   image.Image
	index int
}

func newxpTile(src image.Image) []*xpTile {
	height := src.Bounds().Dy()
	num := height / 1024
	if height > num*1024 {
		num += 1
	}
	res := make([]*xpTile, num)
	for i := 0; i < num; i++ {
		res[i] = &xpTile{src, i}
	}
	return res
}

func (xt *xpTile) At(x, y int) color.Color {
	if x < 256 {
		y1 := y + xt.index*1024
		if y1 < xt.src.Bounds().Dy() {
			return xt.src.At(x, y1)
		}
		return color.RGBA{0, 0, 0, 0}
	} else {
		y1 := y + xt.index*1024 + 512
		if y1 < xt.src.Bounds().Dy() {
			return xt.src.At(x%256, y1)
		}
		return color.RGBA{0, 0, 0, 0}
	}
}

func (xt *xpTile) Bounds() image.Rectangle {
	return image.Rect(0, 0, 512, 512)
}

func (xt *xpTile) ColorModel() color.Model {
	return xt.src.ColorModel()
}
