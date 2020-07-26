package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var content = make([][]uint8, dy)

	for iy, _ := range content {
		content[iy] = make([]uint8, dx)
		for ix, _ := range content[iy] {
			// content[iy][ix] = uint8((iy + ix) / 2)
			// content[iy][ix] = uint8(ix * iy)
			content[iy][ix] = uint8(ix ^ iy)
		}

	}
	return content
}

func main() {
	pic.Show(Pic)
	// pbpaste | base64 -d > /tmp/a.png; open /tmp/a.png
}
