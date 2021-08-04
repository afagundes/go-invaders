package invaders

import (
	"fmt"
	tl "github.com/JoelOtter/termloop"
)

func ShowGameOverScreen(invaders *Invaders) {
	invaders.Level = tl.NewBaseLevel(tl.Cell{Bg: tl.ColorBlack, Fg: tl.ColorWhite})
	invaders.Game.Screen().SetLevel(invaders.Level)

	invaders.initArena()
	invaders.initHud()

	showGameOver(invaders)
	showScore(invaders)
}

func showGameOver(invaders *Invaders) {
	gameOverCanvas := CreateCanvas("invaders/files/game_over.txt")

	arenaX, arenaY := invaders.Arena.Position()
	arenaW, arenaH := invaders.Arena.Size()

	x := arenaX + arenaW/2 - len(gameOverCanvas)/2
	y := arenaY + arenaH/2 + -len(gameOverCanvas[0]) - 1

	invaders.Level.AddEntity(tl.NewEntityFromCanvas(x, y, gameOverCanvas))
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
