package invaders

import tl "github.com/JoelOtter/termloop"

type Alien struct {
	*tl.Entity
}

type AlienType string

const (
	Basic AlienType = "invaders/files/alien_basic.txt"
	Medium AlienType = "invaders/files/alien_medium.txt"
	Strong AlienType = "invaders/files/alien_strong.txt"
)

func NewAlien(alienType AlienType) *Alien {
	canvas := CreateCanvas(string(alienType))
	return &Alien{tl.NewEntityFromCanvas(0, 0, canvas)}
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
