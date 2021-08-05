package invaders

import (
	tl "github.com/JoelOtter/termloop"
)

func CreateCanvas(fileContent []byte) tl.Canvas {
	canvas := tl.CanvasFromString(string(fileContent))
	return canvas
}
