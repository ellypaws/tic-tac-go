package ai

import (
	"tic-tac-toe/board"
)

type PerfectMover struct{}

var PerfectMoverInstance *PerfectMover

func GetPerfectMover() *PerfectMover {
	if PerfectMoverInstance != nil {
		return PerfectMoverInstance
	}
	return &PerfectMover{}
}

func (m *PerfectMover) Move(b *board.Board) (int, int) {
	for i := 0; i < 3; i++ {
		if b.Cells[i][0] != board.Empty && b.Cells[i][0] == b.Cells[i][1] && b.Cells[i][0] == b.Cells[i][2] {
			return i, 0
		}
		if b.Cells[0][i] != board.Empty && b.Cells[0][i] == b.Cells[1][i] && b.Cells[0][i] == b.Cells[2][i] {
			return 0, i
		}
	}
	if b.Cells[0][0] != board.Empty && b.Cells[0][0] == b.Cells[1][1] && b.Cells[0][0] == b.Cells[2][2] {
		return 0, 0
	}
	if b.Cells[0][2] != board.Empty && b.Cells[0][2] == b.Cells[1][1] && b.Cells[0][2] == b.Cells[2][0] {
		return 0, 2
	}
	return GetRandomMover().Move(b)
}
