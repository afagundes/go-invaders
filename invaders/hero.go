package invaders

import (
	_ "embed"
	tl "github.com/JoelOtter/termloop"
	"time"
)

type Hero struct {
	*tl.Entity
	Arena   *Arena
	Lasers  []*Laser
	IsAlive bool
}

//go:embed files/hero.txt
var heroBytes []byte

func NewHero(arena *Arena) *Hero {
	heroCanvas := CreateCanvas(heroBytes)
	x, y := setHeroPosition(arena, heroCanvas)

	return &Hero{Entity: tl.NewEntityFromCanvas(x, y, heroCanvas), Arena: arena, IsAlive: true}
}

func setHeroPosition(arena *Arena, heroCanvas tl.Canvas) (int, int) {
	arenaX, arenaY := arena.Position()
	arenaW, arenaH := arena.Size()

	x := arenaX + arenaW/2 - len(heroCanvas)/2
	y := arenaY + arenaH - len(heroCanvas[0])

	return x, y
}

func (hero *Hero) Tick(event tl.Event) {
	if hero.IsAlive == false {
		return
	}

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
	}
}

func (hero *Hero) shoot() {
	x, y := hero.Position()

	if hero.isReloading() {
		return
	}

	heroWidth, _ := hero.Size()
	heroGunPosition := x + (heroWidth-1)/2
	distanceToHero := y - 1

	laser := NewHeroLaser(heroGunPosition, distanceToHero)
	hero.Lasers = append(hero.Lasers, laser)
}

func (hero *Hero) isReloading() bool {
	return len(hero.Lasers) > 0
}

func (hero *Hero) Collide(collision tl.Physical) {
	if _, ok := collision.(*Laser); ok {
		laser := collision.(*Laser)

		laser.HasHit = true
		hero.IsAlive = false
	}
}

func (hero *Hero) IsDead() bool {
	return hero.IsAlive == false
}

func (hero *Hero) animateHeroEndGame(level *tl.BaseLevel) {
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			level.RemoveEntity(hero)
		} else {
			level.AddEntity(hero)
		}

		time.Sleep(450 * time.Millisecond)
	}
}
