package model

import (
	"encoding/json"
	"model/utils"
)

type Layer struct {
	Inputs      []float64
	InputWidths [][]float64
	// BiasWidths  []float64
	Biases    []float64
	NextLayer *Layer
}

func (l Layer) String() string {
	data, _ := json.Marshal(l)
	return string(data)
}

func (l *Layer) ConnectLayer(nextLayer *Layer) (*Layer, *Layer) {
	inputs := l.Inputs
	nextInputs := nextLayer.Inputs
	inputsSize := len(inputs)
	nextInputSize := len(nextInputs)
	InputWidths := [][]float64{}

	biases := make([]float64, nextInputSize)
	for index, _ := range biases {
		biases[index] = 1
	}
	for l := 0; l < inputsSize; l++ {
		// for range inputs {
		widths := make([]float64, nextInputSize)
		for index, _ := range widths {
			widths[index] = 0.15 // rand.Float64()
		}
		InputWidths = append(InputWidths, widths)
	}

	l.Biases = biases
	l.InputWidths = InputWidths
	l.NextLayer = nextLayer
	// fmt.Println()
	return l, nextLayer
}

func (l *Layer) ForwardPropagation() {
	if l.NextLayer == nil {
		return
	}
	nextLayer := l.NextLayer
	nextLayerInputs := nextLayer.Inputs
	inputWidths := l.InputWidths

	for nextLayerInputIndex, _ := range nextLayerInputs {
		var newInput float64 = 0
		for inputIndex, input := range l.Inputs {

			newInput += inputWidths[inputIndex][nextLayerInputIndex] * input

		}

		nextLayerInputs[nextLayerInputIndex] = utils.Sigmoid(newInput + l.Biases[nextLayerInputIndex])
	}

	//////////////////////////////////////////////////////////
	l.NextLayer.ForwardPropagation()
}

func (l *Layer) BackPropagation(output []float64, learnRate float64) []float64 {

	if l.NextLayer == nil {
		inputs := l.Inputs
		inputsSize := len(inputs)
		if inputsSize != len(output) {
			panic("INPUT SIZE NOT EQUAL OUTPUT SIZE")
		}
		newInputs := make([]float64, inputsSize)
		for index, input := range inputs {
			newInputs[index] = input - output[index]
		}

		l.Inputs = newInputs
		return newInputs
	} else {
		delta := l.NextLayer.BackPropagation(output, learnRate)
		for indexInput, widths := range l.InputWidths {
			for IndexNextInput, width := range widths {

				widths[IndexNextInput] = width + (-learnRate * delta[IndexNextInput] * l.Inputs[indexInput])
			}

		}
		for index, bias := range l.Biases {
			l.Biases[index] = bias + (-learnRate * delta[index])
		}
		newDeltaInput := make([]float64, len(l.Inputs))
		for indexInput, widths := range l.InputWidths {
			var sum float64 = 0
			for IndexNextInput, width := range widths {
				sum += width * delta[IndexNextInput]
			}
			newDeltaInput[indexInput] = sum * utils.SigmoidDerivative(l.Inputs[indexInput])

		}

		l.Inputs = newDeltaInput
	}

	return l.Inputs

}
func (l *Layer) Train(input []float64, output []float64, learnRate float64) {
	inputLayer := NewLayerInput(input)
	inputLayer.ConnectLayer(l)

	l.ForwardPropagation()

	l.BackPropagation(output, learnRate)

}

func NewLayerEmpty(size int) *Layer {
	return &Layer{
		Inputs: make([]float64, size),
	}
}
func NewLayerInput(input []float64) *Layer {
	return &Layer{
		Inputs: input,
	}
}
