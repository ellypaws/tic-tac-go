package neural

import (
	"encoding/gob"
	"math/rand"
	"os"
	"tic-tac-toe/board"
	"time"

	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

type NeuralMover struct {
	model *NeuralNetwork
}

func NewNeuralMover() *NeuralMover {
	return &NeuralMover{
		model: NewNeuralNetwork(),
	}
}

func (m *NeuralMover) Move(b *board.Board) (x int, y int) {
	predictions, _ := m.model.Predict(b)

	// Find the highest probability move that hasn't been made yet
	for {
		maxIdx := 0
		maxVal := -1.0
		for i, p := range predictions {
			if p > maxVal {
				maxVal = p
				maxIdx = i
			}
		}

		x, y = maxIdx/3, maxIdx%3
		if b.Cells[x][y] == board.Empty {
			return x, y
		}

		predictions[maxIdx] = -1.0
	}
}

type NeuralNetwork struct {
	g           *gorgonia.ExprGraph
	input       *gorgonia.Node
	hiddenLayer *gorgonia.Node
	hiddenBias  *gorgonia.Node
	outputLayer *gorgonia.Node
	outputBias  *gorgonia.Node
	output      *gorgonia.Node
	vm          gorgonia.VM
}

func NewNeuralNetwork() *NeuralNetwork {
	g := gorgonia.NewGraph()

	// Input layer
	input := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(1, 20), gorgonia.WithName("input"))

	// Hidden layer. Initialize with small random values
	hiddenLayer := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(20, 10), gorgonia.WithName("hiddenLayer"), gorgonia.WithValue(randomTensor(20, 10)))
	hiddenBias := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(10), gorgonia.WithName("hiddenBias"))
	hidden := gorgonia.Must(gorgonia.Add(gorgonia.Must(gorgonia.Rectify(gorgonia.Must(gorgonia.Mul(input, hiddenLayer)))), hiddenBias))

	// Output layer. Initialize with small random values
	outputLayer := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(10, 9), gorgonia.WithName("outputLayer"), gorgonia.WithValue(randomTensor(10, 9)))
	outputBias := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(9), gorgonia.WithName("outputBias"))
	output := gorgonia.Must(gorgonia.Add(gorgonia.Must(gorgonia.Rectify(gorgonia.Must(gorgonia.Mul(hidden, outputLayer)))), outputBias))

	tapeMachine := gorgonia.NewTapeMachine(g)
	return &NeuralNetwork{
		g:           g,
		input:       input,
		hiddenLayer: hiddenLayer,
		hiddenBias:  hiddenBias,
		outputLayer: outputLayer,
		outputBias:  outputBias,
		output:      output,
		vm:          tapeMachine,
	}
}

func randomTensor(dim1, dim2 int) tensor.Tensor {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	data := make([]float64, dim1*dim2)
	for index := range data {
		data[index] = r.Float64()
	}
	return tensor.New(tensor.WithBacking(data), tensor.WithShape(dim1, dim2))
}

func (nn *NeuralNetwork) Predict(b *board.Board) ([]float64, error) {
	// Convert board to input tensor
	data := make([]float64, 20)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.Cells[i][j] == board.PlayerX {
				data[i*3+j] = 1
			} else if b.Cells[i][j] == board.PlayerO {
				data[9+i*3+j] = 1
			}
		}
	}

	tensorData := tensor.New(tensor.WithShape(1, 20), tensor.Of(tensor.Float64), tensor.WithBacking(data))
	gorgonia.Let(nn.input, tensorData)
	if err := nn.vm.RunAll(); err != nil {
		return nil, err
	}

	// Extract output
	outTensor := nn.output.Value().Data().([]float64)
	return outTensor, nil
}

func boardToInputs(b *board.Board) []float64 {
	data := make([]float64, 20)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			switch b.Cells[i][j] {
			case board.PlayerX:
				data[i*3+j] = 1
			case board.PlayerO:
				data[9+i*3+j] = 1
			}
		}
	}
	return data
}

func (m *NeuralMover) generateDataset() []board.Board {
	// TODO: You can generate dataset by bruteforcing all possible board states
	// for both player's next turn, generate all possible moves
	var dataset []board.Board
	for i := 0; i < 2; i++ {
		for j := 0; j < 9; j++ {
			newBoard := board.NewBoard()
			newBoard.Cells[j/3][j%3] = board.Cell(i)
			dataset = append(dataset, *newBoard)
		}
	}
	return dataset
}

func (m *NeuralMover) Train() {
	// Generate the dataset
	//dataset := m.generateDataset()

	// Train the network using the dataset
	//for _, dataPoint := range dataset {
	// Convert the board to inputs tensor
	//	inputs := boardToInputs(&dataPoint)

	//	 Forward pass
	//	predictions, _ := m.model.Predict(&dataPoint)

	//	 Compute the loss and backprop
	//	 ...
	//}
}

func (m *NeuralMover) Inference(b *board.Board) (x int, y int) {
	//predictions, _ := m.model.Predict(b)

	// code for finding the best move can be similar to the one in the Move() function
	// ...
	return x, y
}

var hiddenLayerWeightsFile = "hiddenLayerWeights.gob"
var outputLayerWeightsFile = "outputLayerWeights.gob"

func (nn *NeuralNetwork) SaveWeights() (hiddenWeightsData []float64, outputWeightsData []float64, err error) {
	hiddenWeightsData = nn.hiddenLayer.Value().Data().([]float64)
	outputWeightsData = nn.outputLayer.Value().Data().([]float64)

	err = saveWeightsToFile(hiddenWeightsData, hiddenLayerWeightsFile)
	if err != nil {
		return nil, nil, err
	}

	err = saveWeightsToFile(outputWeightsData, outputLayerWeightsFile)
	if err != nil {
		return nil, nil, err
	}

	return hiddenWeightsData, outputWeightsData, nil
}

func saveWeightsToFile(weightsData []float64, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(weightsData)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

func (nn *NeuralNetwork) LoadWeights() error {
	hiddenWeightsData, err := loadWeightsFromFile(hiddenLayerWeightsFile)
	if err != nil {
		return err
	}

	outputWeightsData, err := loadWeightsFromFile(outputLayerWeightsFile)
	if err != nil {
		return err
	}

	// Bind the new weights to the nodes
	err = gorgonia.Let(nn.hiddenLayer, tensor.New(tensor.Of(tensor.Float64), tensor.WithShape(20, 10), tensor.WithBacking(hiddenWeightsData)))
	if err != nil {
		return err
	}
	err = gorgonia.Let(nn.outputLayer, tensor.New(tensor.Of(tensor.Float64), tensor.WithShape(10, 9), tensor.WithBacking(outputWeightsData)))
	if err != nil {
		return err
	}

	return nil
}

func loadWeightsFromFile(fileName string) ([]float64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	var weightsData []float64
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&weightsData)
	if err != nil {
		return nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	return weightsData, nil
}
