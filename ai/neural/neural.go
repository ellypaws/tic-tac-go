package neural

import "tic-tac-toe/board"

// from package board
//const (
//	Empty = iota
//	PlayerX
//	PlayerO
//)
//
//type Cell int
//
//type Board struct {
//	Cells [3][3]Cell
//	Id    string
//}

type NeuralMover struct{}

var NeuralMoverInstance *NeuralMover

func GetNeuralMover() *NeuralMover {
	if NeuralMoverInstance != nil {
		return NeuralMoverInstance
	}
	return &NeuralMover{}
}

func (m *NeuralMover) Move(b *board.Board) (x int, y int) {
	// TODO: Implement
	return 0, 0
}

// EVERYTHING BELOW THIS LINE CAN BE COMPLETELY REWRITTEN OR REMOVED

func (m *NeuralMover) generateDataset() []board.Board {
	// TODO: Implement
	return []board.Board{}
}

type NeuralNetwork struct {
}

func (m *NeuralMover) Train() {
	// TODO: Implement
	// I don't know what we should be returning here
	// the weights?
	// use generateDataset() to generate a dataset
	// use move() to make a move within the dataset
}

func move(b board.Board, x int, y int) board.Board {
	// TODO: Implement
	return board.Board{}
}

func (m *NeuralMover) Inference() {
	// TODO: Implement
}
