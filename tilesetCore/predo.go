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
	c := pd.src.At(x3, y3)
	if pd.Hue {
		return pd.mdHue(c)
	}
	return c
}

func max(r, g, b uint8) uint8 {
	if r >= g && r >= b {
		return r
	} else if g >= r && g >= b {
		return g
	}
	return b
}

func min(r, g, b uint8) uint8 {
	if r <= g && r <= b {
		return r
	} else if g <= r && g <= b {
		return g
	}
	return b
}

func center(x uint32) uint8 {
	if x > 255 {
		return 255
	} else if x < 0 {
		return 0
	}
	return uint8(x)
}

func centerFloat32(x float32) float32 {
	if x > 255.0 {
		return 255.0
	} else if x < 0.0 {
		return 0.0
	}
	return x
}

const (
	DHUE   = 0.2
	DLIGHT = 0.1
)

func (pd *PreDo) mdHue(c color.Color) color.Color {
	var (
		S     float32
		alpha float32
	)
	r32, g32, b32, a32 := c.RGBA()
	a := a32 >> 8
	if r32 == g32 && g32 == b32 {
		return c
	}
	if a == 0 {
		return c
	}
	r5, g5, b5 := (r32 / (a + 1)), (g32 / (a + 1)), (b32 / (a + 1))
	r, g, b := center(r5), center(g5), center(b5)
	cMax := max(r, g, b)
	cMin := min(r, g, b)
	delta := (cMax - cMin)
	v := uint16(cMax) + uint16(cMin)

	if v > 255 {
		S = float32(delta) / float32(v)
	} else {
		S = float32(delta) / float32(512-v)
	}

	if (DHUE + S) >= 1.0 {
		alpha = S
	} else {
		alpha = 1.0 - DHUE
	}
	alpha = 1.0 / alpha

	base := float32(v >> 1)
	//fmt.Println(delta, S, base, alpha)
	r1 := base + (float32(r)-base)*alpha
	g1 := base + (float32(g)-base)*alpha
	b1 := base + (float32(b)-base)*alpha

	r1, g1, b1 = centerFloat32(r1), centerFloat32(g1), centerFloat32(b1)

	r2 := r1 * (1.0 - DLIGHT) * float32(a)
	g2 := g1 * (1.0 - DLIGHT) * float32(a)
	b2 := b1 * (1.0 - DLIGHT) * float32(a)

	r3, g3, b3 := uint32(r2), uint32(g2), uint32(b2)

	r4, g4, b4, a4 := uint8(r3>>8), uint8(g3>>8), uint8(b3>>8), uint8(a)
	return color.RGBA{r4, g4, b4, a4}
}

func (pd *PreDo) Bounds() image.Rectangle {
	bd := pd.src.Bounds()
	return image.Rect(0, 0, bd.Dx()*2, bd.Dy()*2)
}

func (pd *PreDo) ColorModel() color.Model {
	return pd.src.ColorModel()
}
