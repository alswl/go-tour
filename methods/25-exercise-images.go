package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	dx, dy int
	data   [][]uint8
}

func NewImage(dx, dy int) *Image {
	var content = make([][]uint8, dy)

	for iy, _ := range content {
		content[iy] = make([]uint8, dx)
		for ix, _ := range content[iy] {
			// content[iy][ix] = uint8((iy + ix) / 2)
			// content[iy][ix] = uint8(ix * iy)
			content[iy][ix] = uint8(ix ^ iy)
		}

	}
	return &Image{dx: dx, dy: dy, data: content}
}

func (i Image) ColorModel() color.Model {
	return color.RGBA64Model
}

func (i Image) Bounds() image.Rectangle {
	return image.Rectangle{image.Pt(0, 0), image.Pt(i.dx, i.dy)}
}

func (i Image) At(x, y int) color.Color {
	point := i.data[x][y]
	return color.RGBA{
		R: point,
		G: point,
		B: point,
		A: 255,
	}
}

func main() {
	m := NewImage(1024, 1024)
	pic.ShowImage(m)
	// pbpaste | base64 -d > /tmp/a.png; open /tmp/a.png
}
