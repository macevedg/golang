package main

import "fmt"

func main() {

	array := []int{-6, -4, 1, 2, 3, 5}

	result := make([]int, len(array))

	leftIndex := 0
	rightIndex := len(array) - 1

	//resultIndex := len(array) - 1

	for i := len(array) - 1; i >= 0; i-- {

		rightNumber := array[rightIndex] * array[rightIndex]
		leftNumber := array[leftIndex] * array[leftIndex]

		if rightNumber < leftNumber {
			result[i] = leftNumber

			leftIndex++
		} else {
			result[i] = rightNumber

			rightIndex--
		}
	}

	fmt.Printf("Result: %d\n", result)
}
