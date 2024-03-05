package main

import (
	"fmt"
)

var (
	valMap = make(map[string]string)
)

func main() {
	array := []int{1, 2, 3, 4, 5}

	//i := 1
	combinations := getCombinations(array)

	for _, a := range combinations {
		arr := toString(a)
		fmt.Printf("[%s],", arr)
	}

}

func getCombinations(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	} else {
		var result [][]int
		for i, num := range nums {
			remaining := make([]int, len(nums)-1)
			copy(remaining, nums[:i])
			copy(remaining[i:], nums[i+1:])
			for _, subset := range getCombinations(remaining) {
				result = append(result, append([]int{num}, subset...))
			}
		}
		return result
	}
}

func toString(array []int) string {

	response := ""

	for _, val := range array {
		response += fmt.Sprintf("%d", val)
	}

	return response
}
