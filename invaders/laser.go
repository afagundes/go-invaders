package invaders

import tl "github.com/JoelOtter/termloop"

type Laser struct {
	*tl.Rectangle
	IsNew  bool
	HasHit bool
}

func NewLaser(heroGunPosition int, y int) *Laser {
	return &Laser{Rectangle: tl.NewRectangle(heroGunPosition, y, 1, 1, tl.ColorRed), IsNew: true, HasHit: false}
}
