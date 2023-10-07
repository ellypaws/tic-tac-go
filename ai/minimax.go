package ai

import "tic-tac-toe/board"

type MinimaxMover struct{}

var MinimaxMoverInstance *MinimaxMover

func GetMinimaxMover() *MinimaxMover {
	if MinimaxMoverInstance != nil {
		return MinimaxMoverInstance
	}
	return &MinimaxMover{}
}

func (m *MinimaxMover) Move(b *board.Board) (int, int) {
	bestVal := -1000
	var bestMoveX, bestMoveY int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.Cells[i][j] == board.Empty {
				b.Cells[i][j] = board.PlayerO
				moveVal := miniMax(b, 0, false)
				b.Cells[i][j] = board.Empty
				if moveVal > bestVal {
					bestMoveX = i
					bestMoveY = j
					bestVal = moveVal
				}
			}
		}
	}
	return bestMoveX, bestMoveY
}

// using minimax algorithm with depth
func miniMax(b *board.Board, depth int, isMax bool) int {
	score := Evaluate(b)
	if score == 10 {
		return score
	}
	if score == -10 {
		return score
	}
	if b.IsFull() {
		return 0
	}

	if isMax {
		best := -1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if b.Cells[i][j] == board.Empty {
					b.Cells[i][j] = board.PlayerO
					best = max(best, miniMax(b, depth+1, !isMax))
					b.Cells[i][j] = board.Empty
				}
			}
		}
		return best
	} else {
		best := 1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if b.Cells[i][j] == board.Empty {
					b.Cells[i][j] = board.PlayerX
					best = min(best, miniMax(b, depth+1, !isMax))
					b.Cells[i][j] = board.Empty
				}
			}
		}
		return best
	}
}

func Evaluate(b *board.Board) int {
	for i := 0; i < 3; i++ {
		if b.Cells[i][0] != board.Empty && b.Cells[i][0] == b.Cells[i][1] && b.Cells[i][0] == b.Cells[i][2] {
			if b.Cells[i][0] == board.PlayerX {
				return -10
			} else {
				return 10
			}
		}
		if b.Cells[0][i] != board.Empty && b.Cells[0][i] == b.Cells[1][i] && b.Cells[0][i] == b.Cells[2][i] {
			if b.Cells[0][i] == board.PlayerX {
				return -10
			} else {
				return 10
			}
		}
	}
	if b.Cells[0][0] != board.Empty && b.Cells[0][0] == b.Cells[1][1] && b.Cells[0][0] == b.Cells[2][2] {
		if b.Cells[0][0] == board.PlayerX {
			return -10
		} else {
			return 10
		}
	}
	if b.Cells[0][2] != board.Empty && b.Cells[0][2] == b.Cells[1][1] && b.Cells[0][2] == b.Cells[2][0] {
		if b.Cells[0][2] == board.PlayerX {
			return -10
		} else {
			return 10
		}
	}
	return 0
}
