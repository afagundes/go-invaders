package invaders

import (
	tl "github.com/JoelOtter/termloop"
)

type Alien struct {
	*tl.Entity
	IsAlive    bool
	IsRendered bool
	Points     int
}

type AlienType struct {
	Source string
	Points int
}

var (
	Basic  = AlienType{Source: "invaders/files/alien_basic.txt", Points: 10}
	Medium = AlienType{Source: "invaders/files/alien_medium.txt", Points: 20}
	Strong = AlienType{Source: "invaders/files/alien_strong.txt", Points: 30}
)

func NewAlien(alienType AlienType) *Alien {
	canvas := CreateCanvas(alienType.Source)
	return &Alien{Entity: tl.NewEntityFromCanvas(0, 0, canvas), IsAlive: true, Points: alienType.Points}
}

func CreateAliensLine(alienType AlienType, lineSize int) []*Alien {
	aliens := make([]*Alien, lineSize)
	for i := 0; i < lineSize; i++ {
		aliens[i] = NewAlien(alienType)
	}

	return aliens
}

func SetPositionAndRenderAliens(aliens [][]*Alien, level *tl.BaseLevel, arena *Arena) {
	initialX, initialY, space := calcInitialPositionAndSpace(aliens, arena)

	for index, line := range aliens {
		x := initialX

		for _, alien := range line {
			_, height := alien.Size()
			y := initialY + height*(index+1) - 2

			alien.SetPosition(x, y)
			alien.IsRendered = true

			level.AddEntity(alien)

			x += space
		}
	}
}

func calcInitialPositionAndSpace(aliens [][]*Alien, arena *Arena) (int, int, int) {
	lineSize := len(aliens[0])
	alienW, _ := aliens[0][0].Size()
	space := alienW + 1

	arenaX, arenaY := arena.Position()
	arenaW, _ := arena.Size()

	totalWidth := lineSize * space
	x := arenaX + arenaW/2 - totalWidth/2

	return x, arenaY, space
}

func (alien *Alien) Collide(collision tl.Physical) {
	if _, ok := collision.(*Laser); ok {
		laser := collision.(*Laser)

		if laser.IsFromHero {
			laser.HasHit = true
			alien.IsAlive = false
		}
	}
}
