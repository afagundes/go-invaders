package invaders

import tl "github.com/JoelOtter/termloop"

type Laser struct {
	*tl.Rectangle
	Direction     int
	IsFromHero    bool
	IsNew         bool
	HasHit        bool
	HitAlienLaser bool
	Points        int
}

func NewHeroLaser(heroGunPosition int, y int) *Laser {
	return &Laser{
		Rectangle:  tl.NewRectangle(heroGunPosition, y, 1, 1, tl.ColorRed),
		Direction:  1,
		IsNew:      true,
		IsFromHero: true,
	}
}

func NewAlienLaser(alienGunPosition int, y int) *Laser {
	return &Laser{
		Rectangle:  tl.NewRectangle(alienGunPosition, y, 1, 1, tl.ColorGreen),
		Direction:  -1,
		IsNew:      true,
		IsFromHero: false,
		Points:     5,
	}
}

func (laser *Laser) Collide(collision tl.Physical) {
	if laser.IsFromHero == false {
		return
	}

	if laserCollide, isLaser := collision.(*Laser); isLaser {
		laser.HasHit = true
		laser.HitAlienLaser = true
		laser.Points = laserCollide.Points

		laserCollide.HasHit = true
	}
}
