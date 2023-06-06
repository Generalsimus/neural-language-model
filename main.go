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
	net := model.CreateModel([]int{
		3,
		5,
		2,
	}, 0.01)

	// net := inputLayer
	// net.ForwardPropagation()
	// fmt.Println("SS1: ", net)
	// net.BackPropagation([]float64{1, 2}, 0.01)
	// input := []float64{1, 2, 3}
	// output := []float64{0, 1}
	fmt.Println("SS2: ", net)
	net.Train([]float64{1, 2, 3}, []float64{0, 1})
	net.Train([]float64{1, 5, 3}, []float64{1, 0})
	fmt.Println("Forward1: ", net.Forward([]float64{1, 2, 3}))
	fmt.Println("Forward2: ", net.Forward([]float64{1, 5, 3}))
	// fmt.Println("Forward2: ", net.Forward(input))
	fmt.Println("SS2: ", net)

	// fmt.Println("SS3: ", net.Forward(input))

	// model.CreateNetworkModel([]int{
	// 	10,
	// 	10,
	// })

}
