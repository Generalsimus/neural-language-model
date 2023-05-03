package model

import (
	"encoding/json"
	"fmt"
	"model/utils"
)

type Network struct {
	Layers    []Layer
	learnRate float64
	// inputs []float64
}

func (n *Network) Forward(inputs []float64) {
	// n.Layers = append(n.Layers, layer)
}

func (n *Network) AddLayer(layer Layer) {
	n.Layers = append(n.Layers, layer)
}

func (n Network) String() string {
	b, err := json.Marshal(n)

	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}

func NewNetwork() Network {
	return Network{
		Layers:    []Layer{},
		learnRate: 0.1,
	}
}

type InputNetwork struct {
	inputs    []Input
	learnRate float64
}

func (network InputNetwork) Train(inputsNum []float64, targetsNum []float64) {
	inputs := network.NumToInputs(inputsNum)
	targets := network.NumToInputs(targetsNum)
	fmt.Println("INPUT: ", inputs, "TARGET: ", targets)

}

func (network InputNetwork) NumToInputs(inputsNum []float64) []Input {
	inputs := make([]Input, len(inputsNum))

	for index, input := range network.inputs {
		value := utils.Find(inputsNum, input.Value)
		if value == input.Value {
			inputs[index] = Input{
				Value:    value,
				Connects: []Connect{},
			}
		} else {
			inputs[index] = input
		}
	}

	return inputs
}
