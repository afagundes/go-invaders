package invaders

import (
	"fmt"
	tl "github.com/JoelOtter/termloop"
)

type Hud struct {
	Title string
	Score *tl.FpsText
}

func NewHud(arena *Arena, level *tl.BaseLevel) *Hud {
	hud := Hud{Title: "Go Invaders"}
	hud.drawTitle(arena, level)
	hud.drawScore(arena, level)

	return &hud
}

func (hud *Hud) drawTitle(arena *Arena, level *tl.BaseLevel) {
	arenaX, arenaY := arena.Position()
	x := arenaX + 1
	y := arenaY - 1

	title := tl.NewText(x, y, hud.Title, tl.ColorWhite, tl.ColorBlack)
	level.AddEntity(title)
}

func (hud *Hud) drawScore(arena *Arena, level *tl.BaseLevel) {
	arenaX, arenaY := arena.Position()
	arenaW, _ := arena.Size()

	txtScore := hud.getScoreText(0)

	x := arenaX + arenaW - len(txtScore) - 1
	y := arenaY - 1

	hud.Score = tl.NewFpsText(x, y, tl.ColorWhite, tl.ColorBlack, 60)
	level.AddEntity(hud.Score)
}

func (hud *Hud) UpdateScore(score int) {
	hud.Score.SetText(hud.getScoreText(score))
}

func (hud *Hud) getScoreText(score int) string {
	txtScore := fmt.Sprintf("Score: %4d", score)
	return txtScore
}
