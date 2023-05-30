package model

import (
	"encoding/json"
	"fmt"
)

type Layer struct {
	Inputs     []float64
	InputWidth [][]float64
	NextLayer  *Layer
}

func (l *Layer) String(prevLayer *Layer) string {
	data, _ := json.Marshal(l)
	return string(data)
}
func (l *Layer) ConnectLayer(prevLayer *Layer) {
	inputs := l.Inputs
	prevInputs := l.Inputs
	inputsSize := len(inputs)
	prevInputSize := len(prevInputs)
	inputWidth := [][]float64{}

	for l := 0; l < inputsSize; l++ {
		widths := make([]float64, prevInputSize)

		inputWidth = append(inputWidth, widths)
	}

	l.InputWidth = inputWidth
	l.NextLayer = prevLayer
	fmt.Println()
}

func NewLayer(size int) *Layer {
	return &Layer{
		Inputs: make([]float64, size),
	}
}
