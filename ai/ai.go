package ai

import (
	"github.com/charmbracelet/bubbles/spinner"
	"math/rand"
	"tic-tac-toe/board"
)

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
	Max
)

type AI struct {
	Difficulty Difficulty
	spinner    spinner.Model
}

func NewAI(difficulty Difficulty) *AI {
	return &AI{Difficulty: difficulty, spinner: spinner.New(spinner.WithSpinner(spinner.Dot))}
}

func (a *AI) Move(b *board.Board) (int, int) {
	var perfectPlayPercentage int

	switch a.Difficulty {
	case Max:
		a.spinner.Update(a.spinner.Tick())
	case Medium:
		perfectPlayPercentage = 20
	case Hard:
		perfectPlayPercentage = 80
	default:
		perfectPlayPercentage = 0
	}

	if rand.Intn(100) < perfectPlayPercentage {
		return PerfectPlay(b)
	}

	return RandomPlay(b)
}

func PerfectPlay(b *board.Board) (int, int) {
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
	return RandomPlay(b)
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

func miniMaxNextMove(b *board.Board) (int, int) {
	bestVal := -1000
	var row, col int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.Cells[i][j] == board.Empty {
				b.Cells[i][j] = board.PlayerO
				moveVal := miniMax(b, 0, false)
				b.Cells[i][j] = board.Empty
				if moveVal > bestVal {
					row = i
					col = j
					bestVal = moveVal
				}
			}
		}
	}
	return row, col
}

func RandomPlay(b *board.Board) (int, int) {
	for {
		row := rand.Intn(3)
		col := rand.Intn(3)
		if b.Cells[row][col] == board.Empty {
			return row, col
		}
	}
}
