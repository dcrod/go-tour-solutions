package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

// TODO: remove unnecessary parts of structure
// since At generates color dynamically from x,y

type Image struct{
	Pix []uint8
	Stride int
	Rect image.Rectangle
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	return i.Rect
}

func (i *Image) At(x, y int) color.Color {
	v := (255.0 / float64(
		(i.Rect.Max.X-i.Rect.Min.X) + (i.Rect.Max.Y-i.Rect.Min.Y)))
	return color.RGBA{0, 0, uint8(v * float64(x+y)), 255}
}

func newImage(r image.Rectangle) *Image {
	return &Image{
		Pix:    make([]uint8, r.Dx() * r.Dy() * 4),
		Stride: r.Dx() * 4,
		Rect:   r,
	}
}

func main() {
	m := newImage(image.Rect(0, 0, 300, 300))
	pic.ShowImage(m)
}
