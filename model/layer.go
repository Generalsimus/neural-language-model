package model

import (
	"encoding/json"
	"fmt"
	"model/utils"
)

type Layer struct {
	// InputSize   int
	LayerWidths [][]float64
	Biases      []float64
	NextLayer   *Layer
}

func (l Layer) String() string {
	data, _ := json.Marshal(l)
	return string(data)
}

func (l *Layer) Fill(inputSize int, outPutSize int) {
	fmt.Println("SIZE: ", inputSize, outPutSize)
	//////////////////////////////
	layerWidths := [][]float64{}
	for i := 0; i < outPutSize; i++ {
		inputWidths := []float64{}
		for i := 0; i < inputSize; i++ {
			inputWidths = append(inputWidths, 0.5)
		}

		layerWidths = append(layerWidths, inputWidths)
	}
	// fmt.Println("layerWidths: ", layerWidths)
	//////////////////////////////
	//////////////////////////////
	biases := make([]float64, outPutSize)
	for i := 0; i < outPutSize; i++ {
		biases[i] = 1
	}
	//////////////////////////////
	l.LayerWidths = layerWidths
	l.Biases = biases
	//////////////////////////////
}

func (l *Layer) Forward(inputs []float64) []float64 {
	layerWidths := l.LayerWidths

	nextLayerInputs := make([]float64, len(layerWidths))
	for layerWidthsIndex, inputWidths := range layerWidths {
		var newInput float64 = 0
		for index, inputWidth := range inputWidths {
			newInput += inputWidth * inputs[index]
		}
		nextLayerInputs[layerWidthsIndex] = utils.Sigmoid(newInput)
		// nextLayerInputs[layerWidthsIndex] = utils.Sigmoid(newInput + l.Biases[layerWidthsIndex])
	}
	// fmt.Println("FORWARD: ", nextLayerInputs)
	if l.NextLayer == nil {
		return nextLayerInputs
	}

	return l.NextLayer.Forward(nextLayerInputs)
}

func (l *Layer) Train(inputs []float64, desiredOutputs []float64, learnRate float64) []float64 {
	// FORWARD ////////////////////////////////////////////
	layerWidths := l.LayerWidths

	nextLayerInputs := make([]float64, len(layerWidths))
	for layerWidthsIndex, inputWidths := range layerWidths {
		var newInput float64 = 0
		for index, inputWidth := range inputWidths {
			newInput += inputWidth * inputs[index]
		}
		nextLayerInputs[layerWidthsIndex] = utils.Sigmoid(newInput)
		// nextLayerInputs[layerWidthsIndex] = utils.Sigmoid(newInput + l.Biases[layerWidthsIndex])
	}
	// fmt.Println("FORWARD: ", nextLayerInputs)

	// BACK PROPAGATION ///////////////////////////////////
	var nextLayerDeltaOutput []float64
	if l.NextLayer == nil {
		nextLayerDeltaOutput = make([]float64, len(nextLayerInputs))
		for index, desiredOutput := range desiredOutputs {
			inputValue := nextLayerInputs[index]
			nextLayerDeltaOutput[index] = inputValue - desiredOutput
			// nextLayerDeltaOutput[index] = utils.Derivative(inputValue) * (desiredOutput - inputValue)
		}
	} else {
		nextLayerDeltaOutput = l.NextLayer.Train(nextLayerInputs, desiredOutputs, learnRate)
	}

	///////////////////////////////////////////////////////
	deltaOutput := make([]float64, len(inputs))

	for inputIndex, input := range inputs {
		var deltaInput float64 = 0
		derivativeInput := utils.Derivative(input)

		for nextLayerDeltaIndex, nextLayerDelta := range nextLayerDeltaOutput {
			width := l.LayerWidths[nextLayerDeltaIndex][inputIndex]

			deltaInput += derivativeInput * (width * nextLayerDelta)

			// fmt.Println("widthsIndex: ", widthsIndex, " inputIndex: ", inputIndex)
			l.LayerWidths[nextLayerDeltaIndex][inputIndex] = (learnRate * nextLayerDelta * input) + width
		}
		deltaOutput[inputIndex] = deltaInput
	}

	return deltaOutput
}
