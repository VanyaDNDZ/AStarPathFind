package astart

import (
	"image"
	"image/color"
	"image/gif"
	"io"
)

const NFrames = 64
const Delay = 20

var Palette = []color.Color{
	color.Black,                        // obstacle
	color.White,                        // Posible, not visited
	color.RGBA{0xFF, 0x00, 0x00, 0xFF}, // visited
	color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // start
	color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // end
	color.RGBA{0xFF, 0x77, 0x77, 0xFF}, // Path
}

func SaveGif(out io.Writer, anim *gif.GIF) error {
	gif.EncodeAll(out, anim)
	return nil
}

func AddFrame(anim *gif.GIF, graph *Graph2d) {
	const (
		size      = 600
		endgeSize = 200
	)

	rect := image.Rect(0, 0, 2*size+1, 2*size+1)
	img := image.NewPaletted(rect, Palette)
	var curPalette uint8
	for i, row := range *graph {
		for j, node := range row {
			if node.NodeType == "o" {
				curPalette = 0
			} else if node.NodeType == "s" {
				curPalette = 3
			} else if node.NodeType == "e" {
				curPalette = 4
			} else if node.NodeType == "v" && node.Visited && !node.InRoute {
				curPalette = 2
			} else if node.NodeType == "v" && node.Visited && node.InRoute {
				curPalette = 5
			} else {
				curPalette = 1
			}
			for x := i * endgeSize; x < (i+1)*endgeSize; x++ {
				for y := j * endgeSize; y < (j+1)*endgeSize; y++ {
					img.SetColorIndex(y, x, curPalette)

				}
			}
		}
	}

	anim.Delay = append(anim.Delay, Delay)
	anim.Image = append(anim.Image, img)

}
