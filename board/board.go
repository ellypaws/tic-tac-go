package board

import (
	"github.com/lrstanley/bubblezone"
)

const (
	Empty = iota
	PlayerX
	PlayerO
)

type Cell int

type Board struct {
	Cells [3][3]Cell
	Id    string
}

func NewBoard() *Board {
	return &Board{
		Id: zone.NewPrefix(),
	}
}

func (b *Board) IsFull() bool {
	for _, row := range b.Cells {
		for _, cell := range row {
			if cell == Empty {
				return false
			}
		}
	}
	return true
}

func (b *Board) Winner() Cell {
	for i := 0; i < 3; i++ {
		if b.Cells[i][0] != Empty && b.Cells[i][0] == b.Cells[i][1] && b.Cells[i][0] == b.Cells[i][2] {
			return b.Cells[i][0]
		}
		if b.Cells[0][i] != Empty && b.Cells[0][i] == b.Cells[1][i] && b.Cells[0][i] == b.Cells[2][i] {
			return b.Cells[0][i]
		}
	}
	if b.Cells[0][0] != Empty && b.Cells[0][0] == b.Cells[1][1] && b.Cells[0][0] == b.Cells[2][2] {
		return b.Cells[0][0]
	}
	if b.Cells[0][2] != Empty && b.Cells[0][2] == b.Cells[1][1] && b.Cells[0][2] == b.Cells[2][0] {
		return b.Cells[0][2]
	}
	return Empty
}
