package ai

import "tic-tac-toe/board"

type NeuralMover struct{}

var NeuralMoverInstance *NeuralMover

func GetNeuralMover() *NeuralMover {
	if NeuralMoverInstance != nil {
		return NeuralMoverInstance
	}
	return &NeuralMover{}
}

func (m *NeuralMover) Move(b *board.Board) (int, int) {
	return 0, 0
}
