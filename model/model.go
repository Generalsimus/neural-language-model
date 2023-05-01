package model

import (
	"encoding/json"
	"fmt"
)

type Network struct {
	Layers []Layer
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
		Layers: []Layer{},
	}
}
