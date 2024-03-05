package main

import (
	"fmt"
)

func main() {

	result := flattened()
	fmt.Printf("Result: %v\n", result)

}

func flattened() []string {
	arr := [][]string{{"a1", "a2", "a3"}, {"b1", "b2", "b3", "b4"}, {"c1"}, {"d1", "d2"}, {"e1"}}

	response := make([]string, 0)

	//0. initialize string array
	//1. iterate over the list of arrays
	for j := 0; j < 4; j++ {
		for i := 0; i < len(arr); i++ {
			// loop
			fmt.Printf("Array: %v\n", response)
			if len(arr[i]) <= j {
				continue
			}
			response = append(response, arr[i][j])
		}

	}

	return response
}

/*
Given the following array

[[a1,a2,a3],[b1,b2,b3,b4],[c1],[d1,d2],[e1]]


Create a function that returns a transpose and flattened version of the array
[a1,b1,c1,d1,e1,a2,b2,d2,a3,b3,b4]
[a1 b1 c1 d1 e1 a2 b2 d2 a3 b3 b4]


Golang code for array declaration
arr:= [][]string{{"a1", "a2", "a3"}, {"b1", "b2", "b3", "b4"}, {"c1"}, {"d1", "d2"}, {"e1"}}
*/
