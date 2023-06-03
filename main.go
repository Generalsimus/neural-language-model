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
	net := model.CreateNetworkModel([]int{
		3,
		2,
	})

	// net := inputLayer
	// net.ForwardPropagation()
	// fmt.Println("SS1: ", net)
	// net.BackPropagation([]float64{1, 2}, 0.01)
	input := []float64{1, 2, 3}
	output := []float64{3, 1}
	fmt.Println("SS2: ", net)
	net.Train(input, output, 0.01)
	fmt.Println("SS3: ", net)
	// model.CreateNetworkModel([]int{
	// 	10,
	// 	10,
	// })

}
