package model

import (
	"fmt"
)

type Layer struct {
	inputs []float64
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
	for index, value := range l.inputs {
		l.inputs[index] = value + num
	}
}

func (l *Layer) AddLayer(layer Layer) {
	for index, value := range l.inputs {
		l.inputs[index] = value + layer.inputs[index]
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

func (l *Layer) MultiplyNum(num float64) {
	for index, value := range l.inputs {
		l.inputs[index] = value * num
	}
}

func (l *Layer) MultiplyLayer(layer Layer) {
	for index, value := range l.inputs {
		l.inputs[index] = value * layer.inputs[index]
	}
}
