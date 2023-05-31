package main

import (
	"fmt"
	"model/model"
)

func main() {
	fmt.Println("SS")
	// input := [][]float64{
	// 	[]float64{2, 5},
	// }
	inputLayer := model.NewLayerInput([]float64{2, 5})

	inputLayer.ConnectLayer(model.CreateNetworkModel([]int{
		3,
		2,
	}))

	net := inputLayer
	net.ForwardPropagation()
	net.BackPropagation([]float64{1, 2}, 0.01)
	fmt.Println("SS", net)
	// model.CreateNetworkModel([]int{
	// 	10,
	// 	10,
	// })

}
