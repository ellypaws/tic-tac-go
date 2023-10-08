package ai

import (
	"github.com/charmbracelet/bubbles/spinner"
	"tic-tac-toe/ai/neural"
	"tic-tac-toe/board"
)

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
	Minimax
	Neural
)

var ai *AI

func Get() *AI {
	if ai != nil {
		return ai
	}
	return ai
}

type AI struct {
	Mover
	Difficulty Difficulty
	spinner    spinner.Model
}

type Mover interface {
	Move(b *board.Board) (int, int)
}

func NewAI(difficulty Difficulty) *AI {
	s := spinner.New(spinner.WithSpinner(spinner.Dot))
	switch difficulty {
	case Easy:
		return &AI{Mover: GetRandomMover(), spinner: s}
	case Medium, Hard:
		return &AI{Mover: GetPerfectMover(), spinner: s}
	case Minimax:
		return &AI{Mover: GetMinimaxMover(), spinner: s}
	case Neural:
		return &AI{Mover: neural.GetNeuralMover(), spinner: s}
	default:
		return &AI{Mover: GetRandomMover(), spinner: s}
	}
}

func (a *AI) ChangeDifficulty(difficulty Difficulty) {
	difficulty = min(difficulty, Neural)
	difficulty = max(difficulty, Easy)

	a.Difficulty = difficulty
	a.Mover = NewAI(difficulty).Mover
}
