package invaders

import (
	"fmt"
	tl "github.com/JoelOtter/termloop"
)

func ShowGameOverScreen(invaders *Invaders) {
	prepareScreen(invaders)
	showGameOver(invaders)
	showScore(invaders)
}

func ShowVictoryScreen(invaders *Invaders) {
	prepareScreen(invaders)
	showVictory(invaders)
	showScore(invaders)
}

func prepareScreen(invaders *Invaders) {
	invaders.Level = tl.NewBaseLevel(tl.Cell{Bg: tl.ColorBlack, Fg: tl.ColorWhite})
	invaders.Game.Screen().SetLevel(invaders.Level)

	invaders.initArena()
	invaders.initHud()
}

func showGameOver(invaders *Invaders) {
	filename := "invaders/files/game_over.txt"
	showCanvas(invaders, filename)
}

func showVictory(invaders *Invaders) {
	filename := "invaders/files/victory.txt"
	showCanvas(invaders, filename)
}

func showCanvas(invaders *Invaders, filename string) {
	canvas := CreateCanvas(filename)

	arenaX, arenaY := invaders.Arena.Position()
	arenaW, arenaH := invaders.Arena.Size()

	x := arenaX + arenaW/2 - len(canvas)/2
	y := arenaY + arenaH/2 + -len(canvas[0]) - 1

	invaders.Level.AddEntity(tl.NewEntityFromCanvas(x, y, canvas))
}

func showScore(invaders *Invaders) {
	score := fmt.Sprintf("SCORE: %4d ", invaders.Score)
	showCenterText(score, 0, invaders)
}

func showCenterText(text string, topPadding int, invaders *Invaders) {
	arenaX, arenaY := invaders.Arena.Position()
	arenaW, arenaH := invaders.Arena.Size()

	x := arenaX + arenaW/2 - len(text)/2
	y := arenaY + arenaH/2 + topPadding

	invaders.Level.AddEntity(tl.NewText(x, y, text, tl.ColorWhite, tl.ColorBlack))
}
