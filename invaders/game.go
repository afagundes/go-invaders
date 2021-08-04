package invaders

import (
	tl "github.com/JoelOtter/termloop"
	"github.com/afagundes/go-invaders/invaders/utils"
	"time"
)

type Invaders struct {
	*tl.Entity
	Game               *tl.Game
	Level              *tl.BaseLevel
	Arena              *Arena
	Hud                *Hud
	Hero               *Hero
	AlienCluster       *AlienCluster
	AlienLaserVelocity float64
	TimeDelta          float64
	RefreshSpeed       time.Duration
	Score              int
}

func NewGame() *Invaders {
	invaders := Invaders{
		Entity:             tl.NewEntity(0, 0, 1, 1),
		Game:               tl.NewGame(),
		Level:              tl.NewBaseLevel(tl.Cell{Bg: tl.ColorBlack, Fg: tl.ColorWhite}),
		AlienLaserVelocity: 0.12,
		RefreshSpeed:       20,
		Score:              0,
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
	invaders.initHud()
	invaders.initHero()
	invaders.initAliens()
	invaders.updatePositions()
}

func (invaders *Invaders) initArena() {
	screenWidth, screenHeight := invaders.getScreenSize()
	invaders.Arena = newArena(screenWidth, screenHeight)
	invaders.Level.AddEntity(invaders.Arena)
}

func (invaders *Invaders) initHud() {
	invaders.Hud = NewHud(invaders.Arena, invaders.Level)
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
	invaders.AlienCluster = NewAlienCluster()
	SetPositionAndRenderAliens(invaders.AlienCluster.Aliens, invaders.Level, invaders.Arena)
}

func (invaders *Invaders) updatePositions() {
	for {
		if invaders.Hero.IsDead() {
			break
		}

		invaders.updateLaserPositions()
		invaders.RemoveDeadAliensAndIncrementScore()
		invaders.updateAlienClusterPosition()

		invaders.Hud.UpdateScore(invaders.Score)

		time.Sleep(invaders.RefreshSpeed * time.Millisecond)
	}
}

func (invaders *Invaders) updateAlienClusterPosition() {
	invaders.AlienCluster.UpdateAliensPositions(invaders.Game.Screen().TimeDelta(), invaders.Arena)
	invaders.AlienCluster.Shoot(invaders.Game.Screen().TimeDelta())
}

func (invaders *Invaders) RemoveDeadAliensAndIncrementScore() {
	points := invaders.AlienCluster.RemoveDeadAliensAndGetPoints(invaders.Level)
	invaders.addScore(points)
}

func (invaders *Invaders) updateLaserPositions() {
	utils.ShowLasersInfo(len(invaders.Hero.Lasers))
	utils.ShowAlienLasersInfo(len(invaders.AlienCluster.Lasers))

	invaders.updateHeroLasers()
	invaders.updateAlienLasers()
	invaders.removeLasers()
}

func (invaders *Invaders) updateHeroLasers() {
	invaders.updateLasers(invaders.Hero.Lasers)
}

func (invaders *Invaders) updateAlienLasers() {
	invaders.TimeDelta += invaders.Game.Screen().TimeDelta()

	if invaders.TimeDelta >= invaders.AlienLaserVelocity {
		invaders.TimeDelta = 0
		invaders.updateLasers(invaders.AlienCluster.Lasers)
	}
}

func (invaders *Invaders) updateLasers(lasers []*Laser) {
	for _, laser := range lasers {
		if laser.IsNew {
			invaders.renderNewLaser(laser)
			continue
		}

		x, y := laser.Position()
		laser.SetPosition(x, y-laser.Direction)
	}
}

func (invaders *Invaders) renderNewLaser(laser *Laser) {
	laser.IsNew = false
	invaders.Level.AddEntity(laser)
}

func (invaders *Invaders) removeLasers() {
	_, arenaY := invaders.Arena.Position()
	_, arenaH := invaders.Arena.Size()

	upperLimit := arenaY
	bottomLimit := arenaY + arenaH - 1

	invaders.Hero.Lasers = invaders.removeLaserOf(invaders.Hero.Lasers, upperLimit)
	invaders.AlienCluster.Lasers = invaders.removeLaserOf(invaders.AlienCluster.Lasers, bottomLimit)
}

func (invaders *Invaders) removeLaserOf(lasers []*Laser, arenaLimit int) []*Laser {
	for index, laser := range lasers {
		_, y := laser.Position()
		isEndOfArena := y == arenaLimit

		if isEndOfArena || laser.HasHit {
			invaders.Level.RemoveEntity(laser)
			lasers = append(lasers[:index], lasers[index+1:]...)
		}
	}

	return lasers
}

func (invaders *Invaders) addScore(points int) {
	invaders.Score += points
}
