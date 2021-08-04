package utils

import (
	"fmt"
	tl "github.com/JoelOtter/termloop"
)

var (
	screenInfo      *tl.FpsText
	arenaInfo       *tl.FpsText
	heroInfo        *tl.FpsText
	lasersInfo      *tl.FpsText
	alienLasersInfo *tl.FpsText
	aliensInfo      *tl.FpsText
	genericInfo     *tl.FpsText
)

func InitInfo(level *tl.BaseLevel) {
	screenInfo = tl.NewFpsText(2, 1, tl.ColorWhite, tl.ColorBlack, 60)
	arenaInfo = tl.NewFpsText(2, 2, tl.ColorWhite, tl.ColorBlack, 60)
	heroInfo = tl.NewFpsText(2, 3, tl.ColorWhite, tl.ColorBlack, 60)
	lasersInfo = tl.NewFpsText(2, 5, tl.ColorWhite, tl.ColorBlack, 60)
	alienLasersInfo = tl.NewFpsText(2, 6, tl.ColorWhite, tl.ColorBlack, 60)
	aliensInfo = tl.NewFpsText(2, 7, tl.ColorWhite, tl.ColorBlack, 60)
	genericInfo = tl.NewFpsText(2, 9, tl.ColorWhite, tl.ColorBlack, 60)

	level.AddEntity(screenInfo)
	level.AddEntity(arenaInfo)
	level.AddEntity(heroInfo)
	level.AddEntity(lasersInfo)
	level.AddEntity(alienLasersInfo)
	level.AddEntity(aliensInfo)
	level.AddEntity(genericInfo)
}

func ShowArenaInfo(screenWidth int, screenHeight int, width int, height int) {
	screenInfo.SetText(formatSizeText("Screen", screenWidth, screenHeight))
	arenaInfo.SetText(formatSizeText("Arena", width, height))
}

func ShowHeroInfo(x, y int) {
	heroInfo.SetText(formatCoordsText("Hero", x, y))
}

func ShowLasersInfo(qtdLasers int) {
	lasersInfo.SetText(fmt.Sprintf("Lasers: %d", qtdLasers))
}

func ShowAlienLasersInfo(qtdLasers int) {
	alienLasersInfo.SetText(fmt.Sprintf("Alien Lasers: %d", qtdLasers))
}

func ShowAliensInfo(qtdAliens int) {
	aliensInfo.SetText(fmt.Sprintf("Aliens: %d", qtdAliens))
}

func ShowGenericInfo(info string) {
	genericInfo.SetText(info)
}

func formatSizeText(text string, x int, y int) string {
	return fmt.Sprintf("%-6s w:%-3d h:%d", text, x, y)
}

func formatCoordsText(text string, x int, y int) string {
	return fmt.Sprintf("%-6s x:%-3d y:%d", text, x, y)
}
