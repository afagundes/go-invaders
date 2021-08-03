package invaders

import (
	"github.com/afagundes/go-invaders/invaders/utils"
)

type AlienCluster struct {
	Aliens                   [][]*Alien
	TimeToMove               float64
	WaitingTime              float64
	WaitingTimeToMoveNextRow float64
	CurrentRowMoving         int
	Direction                int
	MoveSize                 int
	IsMoving                 bool
	ReachedEndArena          bool
	IsMovingDown             bool
}

func NewAlienCluster() *AlienCluster {
	return &AlienCluster{
		TimeToMove:               0,
		WaitingTime:              1,
		WaitingTimeToMoveNextRow: 0.07,
		CurrentRowMoving:         -1,
		Direction:                3,
		MoveSize:                 3,
		IsMoving:                 false,
		ReachedEndArena:          false,
		IsMovingDown:             false,
	}
}

func (alienCluster *AlienCluster) UpdateAliensPositions(timeDelta float64, arena *Arena) {
	utils.ShowAliensInfo(len(alienCluster.Aliens) * len(alienCluster.Aliens[0]))

	alienCluster.TimeToMove += timeDelta

	if alienCluster.CurrentRowMoving == -1 {
		alienCluster.CurrentRowMoving = len(alienCluster.Aliens) - 1
		alienCluster.IsMoving = false
	}

	if alienCluster.ReachedEndArena {
		alienCluster.changeDirection()
	}

	if alienCluster.isWaitingToMove() {
		return
	}

	alienCluster.TimeToMove = 0
	alienCluster.IsMoving = true

	row := alienCluster.Aliens[alienCluster.CurrentRowMoving]

	if alienCluster.IsMovingDown {
		alienCluster.moveDown(row)
	} else {
		alienCluster.moveSideways(row, arena)
	}

	alienCluster.CurrentRowMoving -= 1
}

func (alienCluster *AlienCluster) moveDown(row []*Alien) {
	for _, alien := range row {
		x, y := alien.Position()
		alien.SetPosition(x, y+alienCluster.MoveSize)
	}

	finishedMovingRows := alienCluster.CurrentRowMoving == 0
	if finishedMovingRows {
		alienCluster.IsMovingDown = false
	}
}

func (alienCluster *AlienCluster) moveSideways(row []*Alien, arena *Arena) {
	for _, alien := range row {
		x, y := alien.Position()
		w, _ := alien.Size()
		x = x + alienCluster.Direction

		alien.SetPosition(x, y)
		alienCluster.checkIfReachedEndOfArena(x, w, arena)
	}
}

func (alienCluster *AlienCluster) changeDirection() {
	alienCluster.Direction *= -1
	alienCluster.ReachedEndArena = false
}

func (alienCluster *AlienCluster) isWaitingToMove() bool {
	if alienCluster.isWaitingToMoveFirstRow() || alienCluster.isWaitingToMoveNextRow() {
		return true
	}
	return false
}

func (alienCluster *AlienCluster) isWaitingToMoveFirstRow() bool {
	return alienCluster.IsMoving == false && alienCluster.TimeToMove < alienCluster.WaitingTime
}

func (alienCluster *AlienCluster) isWaitingToMoveNextRow() bool {
	return alienCluster.IsMoving && alienCluster.TimeToMove < alienCluster.WaitingTimeToMoveNextRow
}

func (alienCluster *AlienCluster) checkIfReachedEndOfArena(x int, w int, arena *Arena) {
	if alienCluster.CurrentRowMoving == 0 && (x+w >= arena.End-alienCluster.MoveSize || x <= arena.Init+alienCluster.MoveSize) {
		alienCluster.ReachedEndArena = true
		alienCluster.IsMovingDown = true
	}
}
