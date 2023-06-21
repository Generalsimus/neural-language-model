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
			inputWidths = append(inputWidths, 1)
			// inputWidths = append(inputWidths, rand.Float64())
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
		// nextLayerInputs[layerWidthsIndex] = utils.Sigmoid(newInput)
		nextLayerInputs[layerWidthsIndex] = utils.Sigmoid(newInput + l.Biases[layerWidthsIndex])
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
		// nextLayerInputs[layerWidthsIndex] = utils.Sigmoid(newInput)
		nextLayerInputs[layerWidthsIndex] = utils.Sigmoid(newInput + l.Biases[layerWidthsIndex])
	}
	// fmt.Println("FORWARD: ", nextLayerInputs)

	// BACK PROPAGATION ///////////////////////////////////
	var nextLayerDeltaOutput []float64
	if l.NextLayer == nil {
		nextLayerDeltaOutput = make([]float64, len(nextLayerInputs))
		// errorCost := 0.00
		for index, desiredOutput := range desiredOutputs {
			nextInput := nextLayerInputs[index]
			// errorCost += (desiredOutput - nextInput)

			nextLayerDeltaOutput[index] = (desiredOutput - nextInput) * utils.Derivative(nextInput)
		}
		// fmt.Println("errorCost: ", errorCost)
	} else {
		nextLayerDeltaOutput = l.NextLayer.Train(nextLayerInputs, desiredOutputs, learnRate)
	}

	///////////////////////////////////////////////////////
	deltaOutput := make([]float64, len(inputs))
	// momentum := 0.00

	for inputIndex, input := range inputs {
		delta := 0.00
		for widthsIndex, widths := range l.LayerWidths {
			nextLayerDelta := nextLayerDeltaOutput[widthsIndex]
			width := widths[inputIndex]

			// error += outputNode.Error * outputNode.Input[node].Weight * (node.Output * (1.0 - node.Output))
			delta += (nextLayerDelta * width)

			// outputNode.Input[node].Weight += m_learningRate * outputNode.Error * node.Output;
			widths[inputIndex] += learnRate * nextLayerDelta * input
			// nextLayerInputs[widthsIndex]

			// neuron.bias.weight += neuron.bias.delta * learningRate;
			// m_learningRate * outputNode.Error * outputNode.Bias.Weight
			// l.Biases[widthsIndex] += nextLayerDelta * learnRate
			l.Biases[widthsIndex] += nextLayerDelta * learnRate
			// l.Biases[widthsIndex] += (nextLayerDelta * l.Biases[widthsIndex]) * learnRate
		}

		deltaOutput[inputIndex] = delta * utils.Derivative(input)
	}

	return deltaOutput
}
