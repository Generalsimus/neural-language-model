package model

import (
	"encoding/json"
	"fmt"
	"model/utils"
)

type Layer struct {
	Inputs      []float64
	InputWidths [][]float64
	NextLayer   *Layer
}

func (l Layer) String() string {
	data, _ := json.Marshal(l)
	return string(data)
}

func (l *Layer) ConnectLayer(prevLayer *Layer) (*Layer, *Layer) {
	inputs := l.Inputs
	prevInputs := prevLayer.Inputs
	inputsSize := len(inputs)
	prevInputSize := len(prevInputs)
	InputWidths := [][]float64{}

	for l := 0; l < inputsSize; l++ {
		widths := make([]float64, prevInputSize)
		for index, _ := range widths {
			widths[index] = 0.15 // rand.Float64()
		}
		InputWidths = append(InputWidths, widths)
	}

	l.InputWidths = InputWidths
	l.NextLayer = prevLayer
	fmt.Println()
	return l, prevLayer
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
			// fmt.Println(inputIndex, inputWidths)
			// fmt.Println(nextLayerInputIndex)
			// fmt.Println(inputWidths[inputIndex])
			newInput += inputWidths[inputIndex][nextLayerInputIndex] * input

		}

		nextLayerInputs[nextLayerInputIndex] = utils.Sigmoid(newInput)
	}

	//////////////////////////////////////////////////////////
	l.NextLayer.ForwardPropagation()
}

func (l *Layer) BackPropagation(outputNum []float64, learnRate float64) float64 {
	var delta float64 = 0
	if l.NextLayer != nil {
		delta = l.NextLayer.BackPropagation(outputNum, learnRate)
	} else {
		inputs := l.Inputs
		inputsSize := len(inputs)
		if inputsSize != len(outputNum) {
			panic("INPUT SIZE NOT EQUAL OUTPUT SIZE")
		}
		newInputs := make([]float64, len(inputs))
		for index, input := range inputs {
			newInputs[index] = input - outputNum[index]
		}
		fmt.Println("ELLL", l)
		l.Inputs = newInputs
		fmt.Println("ESSS", l)

	}
	return delta

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
