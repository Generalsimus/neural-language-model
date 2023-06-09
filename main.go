package main

import (
	"fmt"
	"model/model"
	"model/utils"
)

func main() {
	fmt.Println("SS")
	// input := [][]float64{
	// 	[]float64{2, 5},
	// }
	net := model.CreateModel([]int{
		3,
		// 2,
		2,
	}, 0.01)

	for i := 0; i < 5; i++ {

		net.Train([]float64{1, 2, 3}, []float64{0, 1})
		net.Train([]float64{1, 5, 3}, []float64{1, 0})
	}
	fmt.Println(utils.Sigmoid(0.4978321736604644*1 + 0.4956831970068629*2 + 0.4934965209813932*3))

	fmt.Println("SS2: ", net)
	fmt.Println("EEEE: ", utils.Sigmoid(10)*(1-utils.Sigmoid(10)), 0.5*10)

	forward1 := net.Forward([]float64{1, 2, 3})
	m1 := FindMaxIndex(forward1)

	fmt.Println("Forward1: ", forward1, m1, forward1[1] > forward1[0])

	forward2 := net.Forward([]float64{1, 5, 3})
	m2 := FindMaxIndex(forward2)

	fmt.Println("Forward2: ", forward2, m2, forward2[1] > forward2[0])

}
func FindMaxIndex[T int | uint | float64 | int64 | float32](a []T) int {
	maxIndex := 0
	max := a[maxIndex]
	for index, value := range a {

		if value > max {
			max = value
			maxIndex = index
		}
	}
	return maxIndex
}
