package invaders

import (
	_ "embed"
	"fmt"
	tl "github.com/JoelOtter/termloop"
)

var (
	//go:embed files/title.txt
	titleScreenFile []byte

	//go:embed files/game_over.txt
	gameOverScreenFile []byte
)

func ShowTitleScreen(invaders *Invaders) {
	prepareScreen(invaders)
	showTitle(invaders)

	if checkArenaSizeNotOk(invaders) {
		invaders.ScreenSizeNotOK = true
		showMaximizeScreen(invaders)
		return
	}

	showPressToInit(invaders, 0)
}

func checkArenaSizeNotOk(invaders *Invaders) bool {
	w, h := invaders.Arena.Size()

	if w < 100 || h < 37 {
		return true
	}

	return false
}

func ShowGameOverScreen(invaders *Invaders) {
	prepareScreen(invaders)
	showGameOver(invaders)
	showScore(invaders)
	showPressToInit(invaders, 2)
}

func prepareScreen(invaders *Invaders) {
	invaders.Level = tl.NewBaseLevel(tl.Cell{Bg: tl.ColorBlack, Fg: tl.ColorWhite})
	invaders.Game.Screen().SetLevel(invaders.Level)
	invaders.Level.AddEntity(invaders)

	invaders.initArena()
	invaders.initHud()
}

func showTitle(invaders *Invaders) {
	showCanvas(invaders, titleScreenFile)
}

func showGameOver(invaders *Invaders) {
	showCanvas(invaders, gameOverScreenFile)
}

func showCanvas(invaders *Invaders, file []byte) {
	canvas := CreateCanvas(file)

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

func showPressToInit(invaders *Invaders, topPadding int) {
	showCenterText("Press ENTER to start", topPadding, invaders)
}

func showMaximizeScreen(invaders *Invaders) {
	showCenterText("Maximize the console and run the game again", 0, invaders)
}

func showCenterText(text string, topPadding int, invaders *Invaders) {
	arenaX, arenaY := invaders.Arena.Position()
	arenaW, arenaH := invaders.Arena.Size()

	x := arenaX + arenaW/2 - len(text)/2
	y := arenaY + arenaH/2 + topPadding

	invaders.Level.AddEntity(tl.NewText(x, y, text, tl.ColorWhite, tl.ColorBlack))
}
