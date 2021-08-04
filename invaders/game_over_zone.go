package invaders

import tl "github.com/JoelOtter/termloop"

type GameOverZone struct {
	*tl.Entity
	EnteredZone bool
}

func CreateGameOverZone(arena *Arena, hero *Hero) *GameOverZone {
	_, heroH := hero.Size()
	arenaX, arenaY := arena.Position()
	arenaW, arenaH := arena.Size()

	x := arenaX
	y := arenaY + arenaH - heroH
	w := arenaW
	h := heroH

	return &GameOverZone{Entity: tl.NewEntity(x, y, w, h)}
}

func (gameOverZone *GameOverZone) Collide(collision tl.Physical) {
	if _, ok := collision.(*Alien); ok && collision.(*Alien).IsAlive {
		gameOverZone.EnteredZone = true
	}
}
