package invaders

import (
	tl "github.com/JoelOtter/termloop"
	"github.com/afagundes/go-invaders/invaders/utils"
	"time"
)

type Invaders struct {
	*tl.Entity
	Game         *tl.Game
	Level        *tl.BaseLevel
	Arena        *Arena
	Hero         *Hero
	AlienCluster *AlienCluster
}

func NewGame() *Invaders {
	invaders := Invaders{
		Entity: tl.NewEntity(0, 0, 1, 1),
		Game:   tl.NewGame(),
		Level:  tl.NewBaseLevel(tl.Cell{Bg: tl.ColorBlack, Fg: tl.ColorWhite}),
	}

	invaders.Game.Screen().SetFps(60)
	invaders.Game.Screen().SetLevel(invaders.Level)
	invaders.Level.AddEntity(&invaders)

	utils.InitInfo(invaders.Level)

	return &invaders
}

func (invaders *Invaders) Start() {
	go invaders.initializeGame()
	invaders.Game.Start()
}

func (invaders *Invaders) initializeGame() {
	invaders.initArena()
	invaders.initHero()
	invaders.initAliens()
	invaders.updatePositions()
}

func (invaders *Invaders) initArena() {
	screenWidth, screenHeight := invaders.getScreenSize()
	invaders.Arena = newArena(screenWidth, screenHeight)
	invaders.Level.AddEntity(invaders.Arena)
}

func (invaders *Invaders) getScreenSize() (int, int) {
	screenWidth, screenHeight := invaders.Game.Screen().Size()

	for screenWidth == 0 && screenHeight == 0 {
		time.Sleep(100 * time.Millisecond)
		screenWidth, screenHeight = invaders.Game.Screen().Size()
	}

	return screenWidth, screenHeight
}

func (invaders *Invaders) initHero() {
	invaders.Hero = NewHero(invaders.Arena)
	invaders.Level.AddEntity(invaders.Hero)
}

func (invaders *Invaders) initAliens() {
	lineSize := 13

	invaders.AlienCluster = NewAlienCluster()
	invaders.AlienCluster.Aliens = append(invaders.AlienCluster.Aliens, CreateAliensLine(Strong, lineSize))
	invaders.AlienCluster.Aliens = append(invaders.AlienCluster.Aliens, CreateAliensLine(Medium, lineSize))
	invaders.AlienCluster.Aliens = append(invaders.AlienCluster.Aliens, CreateAliensLine(Medium, lineSize))
	invaders.AlienCluster.Aliens = append(invaders.AlienCluster.Aliens, CreateAliensLine(Basic, lineSize))
	invaders.AlienCluster.Aliens = append(invaders.AlienCluster.Aliens, CreateAliensLine(Basic, lineSize))

	SetPositionAndRenderAliens(invaders.AlienCluster.Aliens, invaders.Level, invaders.Arena)
}

func (invaders *Invaders) updatePositions() {
	var refreshSpeed time.Duration = 20

	for {
		invaders.updateLaserPositions()
		invaders.updateAlienClusterPosition()

		time.Sleep(refreshSpeed * time.Millisecond)
	}
}

func (invaders *Invaders) updateLaserPositions() {
	utils.ShowLasersInfo(len(invaders.Hero.Lasers))

	for _, laser := range invaders.Hero.Lasers {
		if laser.IsNew {
			invaders.renderNewLaser(laser)
			continue
		}

		x, y := laser.Position()
		laser.SetPosition(x, y-1)
	}

	invaders.removeLasers()
}

func (invaders *Invaders) updateAlienClusterPosition() {
	invaders.removeDeadAliens()
	invaders.AlienCluster.UpdateAliensPositions(invaders.Game.Screen().TimeDelta(), invaders.Arena)
}

func (invaders *Invaders) removeDeadAliens() {
	for _, alienRow := range invaders.AlienCluster.Aliens {
		for _, alien := range alienRow {
			if alien.IsAlive == false {
				invaders.Level.RemoveEntity(alien)
			}
		}
	}
}

func (invaders *Invaders) renderNewLaser(laser *Laser) {
	laser.IsNew = false
	invaders.Level.AddEntity(laser)
}

func (invaders *Invaders) removeLasers() {
	_, arenaY := invaders.Arena.Position()

	for index, laser := range invaders.Hero.Lasers {
		_, y := laser.Position()
		isEndOfArena := y == arenaY

		if isEndOfArena || laser.HasHit {
			invaders.Level.RemoveEntity(laser)
			invaders.Hero.Lasers = append(invaders.Hero.Lasers[:index], invaders.Hero.Lasers[index+1:]...)
		}
	}
}
