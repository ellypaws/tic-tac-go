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

func New() Model {
	var player board.Cell
	start := rand.Int() % 2
	if start == 0 {
		player = board.PlayerX
	} else {
		player = board.PlayerO
	}
	return Model{
		Board:   board.NewBoard(),
		Current: player,
		help:    help.New(),
		ai:      ai.NewAI(ai.Medium),
	}
}
