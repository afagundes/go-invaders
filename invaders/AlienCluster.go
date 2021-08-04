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
	lineSize := 13

	alienCluster := AlienCluster{
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

	alienCluster.Aliens = append(alienCluster.Aliens, CreateAliensLine(Strong, lineSize))
	alienCluster.Aliens = append(alienCluster.Aliens, CreateAliensLine(Medium, lineSize))
	alienCluster.Aliens = append(alienCluster.Aliens, CreateAliensLine(Medium, lineSize))
	alienCluster.Aliens = append(alienCluster.Aliens, CreateAliensLine(Basic, lineSize))
	alienCluster.Aliens = append(alienCluster.Aliens, CreateAliensLine(Basic, lineSize))

	return &alienCluster
}

func (alienCluster *AlienCluster) UpdateAliensPositions(timeDelta float64, arena *Arena) {
	utils.ShowAliensInfo(len(alienCluster.Aliens) * len(alienCluster.Aliens[0]))

	alienCluster.prepareMovement()

	if alienCluster.canMove(timeDelta) {
		alienCluster.move(arena)
	}
}

func (alienCluster *AlienCluster) prepareMovement() {
	if alienCluster.CurrentRowMoving == -1 {
		alienCluster.CurrentRowMoving = len(alienCluster.Aliens) - 1
		alienCluster.IsMoving = false
	}

	if alienCluster.ReachedEndArena {
		alienCluster.changeDirection()
	}
}

func (alienCluster *AlienCluster) changeDirection() {
	alienCluster.Direction *= -1
	alienCluster.ReachedEndArena = false
}

func (alienCluster *AlienCluster) canMove(timeDelta float64) bool {
	alienCluster.TimeToMove += timeDelta
	if alienCluster.isWaitingToMoveFirstRow() || alienCluster.isWaitingToMoveNextRow() {
		return false
	}
	return true
}

func (alienCluster *AlienCluster) isWaitingToMoveFirstRow() bool {
	return alienCluster.IsMoving == false && alienCluster.TimeToMove < alienCluster.WaitingTime
}

func (alienCluster *AlienCluster) isWaitingToMoveNextRow() bool {
	return alienCluster.IsMoving && alienCluster.TimeToMove < alienCluster.WaitingTimeToMoveNextRow
}

func (alienCluster *AlienCluster) move(arena *Arena) {
	alienCluster.TimeToMove = 0
	alienCluster.IsMoving = true

	if alienCluster.IsMovingDown {
		alienCluster.moveDown()
	} else {
		alienCluster.moveSideways(arena)
	}

	alienCluster.CurrentRowMoving -= 1
}

func (alienCluster *AlienCluster) moveDown() {
	row := alienCluster.getCurrentLine()

	for _, alien := range row {
		x, y := alien.Position()
		alien.SetPosition(x, y+alienCluster.MoveSize)
	}

	finishedMovingRows := alienCluster.CurrentRowMoving == 0
	if finishedMovingRows {
		alienCluster.IsMovingDown = false
	}
}

func (alienCluster *AlienCluster) moveSideways(arena *Arena) {
	row := alienCluster.getCurrentLine()

	for _, alien := range row {
		x, y := alien.Position()
		w, _ := alien.Size()
		x = x + alienCluster.Direction

		alien.SetPosition(x, y)
		alienCluster.checkIfReachedEndOfArena(x, w, arena)
	}
}

func (alienCluster *AlienCluster) getCurrentLine() []*Alien {
	row := alienCluster.Aliens[alienCluster.CurrentRowMoving]
	return row
}

func (alienCluster *AlienCluster) checkIfReachedEndOfArena(x int, w int, arena *Arena) {
	if alienCluster.CurrentRowMoving == 0 && (x+w >= arena.End-alienCluster.MoveSize || x <= arena.Init+alienCluster.MoveSize) {
		alienCluster.ReachedEndArena = true
		alienCluster.IsMovingDown = true
	}
}
