package model

import (
	"encoding/json"
)

// func CreateNetworkModel(layersSize []int) *Layer {
// 	argCount := len(layersSize)
// 	if argCount < 2 {
// 		panic("MINIMAL LAYER SIZE IS 2")
// 	}

// 	layer1 := NewLayerEmpty(layersSize[0])

// 	for i := 1; i < argCount; i++ {
// 		size := layersSize[i]
// 		layer2 := NewLayerEmpty(size)
// 		layer2.i
// 		// layer1.ConnectLayer(layer2)
// 	}

//		return layer1
//	}
type Network struct {
	LearnRate  float64
	StartLayer *Layer
	// EndLayer   *Layer
}

func (n *Network) String() string {
	data, _ := json.Marshal(n)
	return string(data)
}
func (n *Network) Train(inputs []float64, outputs []float64) {

	n.StartLayer.Train(inputs, outputs, n.LearnRate)
}

func (n *Network) Forward(inputs []float64) []float64 {

	return n.StartLayer.Forward(inputs)
}

func CreateModel(layersSize []int, learnRate float64) *Network {
	argCount := len(layersSize)
	if argCount < 2 {
		panic("MINIMAL LAYER SIZE IS 2")
	}
	///////////////////////////////////////////////////////////////
	startLayer := &Layer{}
	previousLayer := startLayer

	///////////////////////////////////////////////////////////////
	for i := 1; i < argCount; i++ {
		previousLayer.Fill(layersSize[i-1], layersSize[i])
		if i != (argCount - 1) {
			newLayer := &Layer{}
			previousLayer.NextLayer = newLayer
			previousLayer = newLayer
		}
	}
	// fmt.Println("MODEL: ", startLayer)
	return &Network{
		LearnRate:  learnRate,
		StartLayer: startLayer,
		// EndLayer:   previousLayer,
	}
}
