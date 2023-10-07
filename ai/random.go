package ai

import (
	"math/rand"
	"tic-tac-toe/board"
)

type RandomMover struct{}

var RandomMoverInstance *RandomMover

func GetRandomMover() *RandomMover {
	if RandomMoverInstance != nil {
		return RandomMoverInstance
	}
	return &RandomMover{}
}

func (m *RandomMover) Move(b *board.Board) (int, int) {
	for {
		row := rand.Intn(3)
		col := rand.Intn(3)
		if b.Cells[row][col] == board.Empty {
			return row, col
		}
	}
}
