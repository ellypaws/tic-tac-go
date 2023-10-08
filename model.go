package main

import (
	"math/rand"
	"tic-tac-toe/ai"
	"tic-tac-toe/board"
	"tic-tac-toe/help"
)

type Model struct {
	Board         *board.Board
	Current       board.Cell
	gameOver      bool
	invalidMove   bool
	Cursor        struct{ Row, Col int }
	Width, Height int
	help          help.Model
	ai            *ai.AI
}

type opts func(*Model)

func New(options ...opts) Model {
	var player board.Cell
	start := rand.Int() % 2
	if start == 0 {
		player = board.PlayerX
	} else {
		player = board.PlayerO
	}
	m := Model{
		Board:    board.NewBoard(),
		Current:  player,
		gameOver: false,
		help:     help.New(),
		ai:       ai.NewAI(ai.Medium),
	}
	for _, option := range options {
		option(&m)
	}
	return m
}

func WithAI(ai *ai.AI) opts {
	return func(m *Model) {
		m.ai = ai
	}
}

func WithDifficulty(difficulty ai.Difficulty) opts {
	return func(m *Model) {
		m.ai.ChangeDifficulty(difficulty)
	}
}

func WithStartPlayer(player board.Cell) opts {
	return func(m *Model) {
		m.Current = player
	}
}

func WithWidth(width int) opts {
	return func(m *Model) {
		m.Width = width
	}
}

func WithHeight(height int) opts {
	return func(m *Model) {
		m.Height = height
	}
}
