package main

import (
	"fmt"
	"model/model"
	"model/utils"
	// "model/utils"
)

func main() {
	// input := 0.85

	// desiredOutput := 0.7

	// error := desiredOutput - input
	// delta := error * utils.Derivative(input)

	// learnRate := 0.01
	// fmt.Println("error: ", error)
	// fmt.Println("Derivative: ", utils.Derivative(input))
	// fmt.Println("learnRate: ", learnRate)
	// fmt.Println("delta: ", delta)
	// fmt.Println("n2: ", learnRate*delta)
	input := 1.5

	width := 0.8

	output := utils.Sigmoid(input * width)
	outputDer := utils.Derivative(output)

	fmt.Println(input, width, output, outputDer, output-outputDer)
	// return
	net := model.CreateModel([]int{
		3,
		16,
		2,
	}, 0.01)
	input1 := []float64{0.1, 0.2, 0.3}
	output1 := []float64{0.1, 0.0}

	input2 := []float64{0.1, 0.5, 0.3}
	output2 := []float64{0, 0.1}

	// for i := 0; i < 400000; i++ {
	// 	// fmt.Println("LEARN1")
	// net.Train(input1, output1)

	// // 	// fmt.Println("LEARN2")
	net.Train(input2, output2)
	// }

	fmt.Println("SS2: ", net)

	forward1 := net.Forward(input1)
	m1 := FindMaxIndex(forward1)

	fmt.Println("Forward1: ", forward1, m1, FindMaxIndex(forward1) == FindMaxIndex(output1), output1)

	forward2 := net.Forward(input2)
	m2 := FindMaxIndex(forward2)

	fmt.Println("Forward2: ", forward2, m2, FindMaxIndex(forward2) == FindMaxIndex(output2), output2)

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
