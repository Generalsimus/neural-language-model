package model

import (
	"fmt"
)

type Layer struct {
	Inputs  []float64
	Weights [][]float32
	Biases  []float32
}

func (l *Layer) AddInputItem(item float64) {
	l.Inputs = append(l.Inputs, item)
	l.Weights = append(l.Weights, []float32{})
	l.Biases = append(l.Biases, 1)
}

func (l *Layer) Add(el interface{}) {
	switch v := el.(type) {
	case Layer:
		l.AddLayer(v)
	case float64:
		l.AddNum(v)
	default:
		panic(fmt.Sprintf("Layer Add Argument Type Error: %d", el))
	}
}

func (l *Layer) AddNum(num float64) {
	for index, value := range l.Inputs {
		l.Inputs[index] = value + num
	}
}

func (l *Layer) AddLayer(layer Layer) {
	for index, value := range l.Inputs {
		l.Inputs[index] = value + layer.Inputs[index]
	}
}

func (l *Layer) Map(callBackGetValue func(v float64, index int) float64) {
	for index, value := range l.Inputs {
		l.Inputs[index] = callBackGetValue(value, index)
	}
}

func (l *Layer) Multiply(el interface{}) {
	switch v := el.(type) {
	case Layer:
		l.MultiplyLayer(v)
	case float64:
		l.MultiplyNum(v)
	default:
		panic(fmt.Sprintf("Layer Multiply Argument Type Error: %d", el))
	}
}

// TODO: fix it
func (l *Layer) Transpose(num float64) Layer {
	inputsCount := len(l.Inputs)
	inputs := make([]float64, inputsCount)
	for index, value := range l.Inputs {
		inputs[inputsCount-index-1] = value
	}
	return Layer{
		Inputs: inputs,
	}
}
func (l *Layer) MultiplyNum(num float64) {
	for index, value := range l.Inputs {
		l.Inputs[index] = value * num
	}
}

func (l *Layer) MultiplyLayer(layer Layer) {
	for index, value := range l.Inputs {
		l.Inputs[index] = value * layer.Inputs[index]
	}
}

func NewLayer() Layer {
	return Layer{
		Inputs:  []float64{},
		Weights: [][]float32{},
		Biases:  []float32{},
	}
}
