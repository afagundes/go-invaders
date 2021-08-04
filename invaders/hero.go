package invaders

import (
	tl "github.com/JoelOtter/termloop"
	"github.com/afagundes/go-invaders/invaders/utils"
)

type Hero struct {
	*tl.Entity
	Arena  *Arena
	Lasers []*Laser
}

func NewHero(arena *Arena) *Hero {
	heroCanvas := CreateCanvas("invaders/files/hero.txt")
	x, y := setHeroPosition(arena, heroCanvas)

	utils.ShowHeroInfo(x, y)

	return &Hero{Entity: tl.NewEntityFromCanvas(x, y, heroCanvas), Arena: arena}
}

func setHeroPosition(arena *Arena, heroCanvas tl.Canvas) (int, int) {
	arenaX, arenaY := arena.Position()
	arenaW, arenaH := arena.Size()

	x := arenaX + arenaW/2 - len(heroCanvas)/2
	y := arenaY + arenaH - len(heroCanvas[0])

	return x, y
}

func (hero *Hero) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		x, y := hero.Position()
		heroWidth, _ := hero.Size()

		switch event.Key {
		case tl.KeyArrowLeft:
			if x > hero.Arena.Init {
				x = x - 1
				hero.SetPosition(x, y)
			}
		case tl.KeyArrowRight:
			if x < hero.Arena.End-heroWidth-1 {
				x = x + 1
				hero.SetPosition(x, y)
			}
		case tl.KeySpace:
			hero.shoot()
		}

		utils.ShowHeroInfo(x, y)
	}
}

func (hero *Hero) shoot() {
	x, y := hero.Position()

	if hero.isReloading(y) {
		return
	}

	heroWidth, _ := hero.Size()
	heroGunPosition := x + (heroWidth-1)/2
	distanceToHero := y - 1

	laser := NewHeroLaser(heroGunPosition, distanceToHero)
	hero.Lasers = append(hero.Lasers, laser)
}

func (hero *Hero) isReloading(y int) bool {
	distanceBetweenShots := 3

	if len(hero.Lasers) > 0 {
		lastLaser := hero.Lasers[len(hero.Lasers)-1]
		if _, lastPosition := lastLaser.Position(); lastPosition >= y-distanceBetweenShots {
			return true
		}
	}

	return false
}
