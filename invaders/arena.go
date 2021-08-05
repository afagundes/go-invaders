package invaders

import (
	tl "github.com/JoelOtter/termloop"
	"github.com/afagundes/go-invaders/invaders/utils"
)

type Arena struct {
	*tl.Entity
	Init int
	End  int
}

const (
	ArenaMaxWidth  = 100
	ArenaMaxHeight = 37
)

func newArena(screenWidth int, screenHeight int) *Arena {
	width := utils.ValueMinusPercent(screenWidth, 0.40)
	height := utils.ValueMinusPercent(screenHeight, 0.15)

	if width > ArenaMaxWidth {
		width = ArenaMaxWidth
	}

	if height > ArenaMaxHeight {
		height = ArenaMaxHeight
	}

	centerX := screenWidth/2 - width/2
	centerY := screenHeight/2 - height/2
	init := centerX + 1
	end := centerX + width

	return &Arena{tl.NewEntityFromCanvas(centerX, centerY, createArena(width, height)), init, end}
}

func createArena(width, height int) tl.Canvas {
	canvas := tl.NewCanvas(width, height)

	for x, cell := range canvas {
		for y := range cell {
			fillTopBottom(x, y, height, canvas)
			fillSides(x, y, width, canvas)
		}
	}

	createCell(0, 0, canvas, '┌')
	createCell(width-1, 0, canvas, '┐')
	createCell(0, height-1, canvas, '└')
	createCell(width-1, height-1, canvas, '┘')

	return canvas
}

func fillTopBottom(x, y, height int, canvas tl.Canvas) {
	if x > 0 && (y == 0 || y == height-1) {
		createCell(x, y, canvas, '─')
	}
}

func fillSides(x, y, width int, canvas tl.Canvas) {
	if x == 0 || x == width-1 {
		createCell(x, y, canvas, '│')
	}
}

func createCell(x, y int, canvas tl.Canvas, ch rune) {
	canvas[x][y] = tl.Cell{Bg: tl.ColorBlack, Fg: tl.ColorWhite, Ch: ch}
}
