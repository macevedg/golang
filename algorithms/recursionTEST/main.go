package main

import "fmt"

func main() {

	number := recursion(4)
	fmt.Printf("Recursive: %d\n", number)
}

func recursion(number int) int {

	if number == 1 {

		return number
	}

	return number + recursion(number-1)
}
