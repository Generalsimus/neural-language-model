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
		10,
		10,
	})
	fmt.Println("SS", net)

}
