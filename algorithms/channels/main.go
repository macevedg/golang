package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	list1 := []int{1, 2, 3, 4, 5}
	list2 := []int{2, 20, 30, 40, 50, 60, -10, -10, -10, 30, 30, 30, 30, 30, 30}

	a := make(chan int)
	b := make(chan int)

	go partialSum(list1, a)
	go partialSum(list2, b)

	x, y := <-a, <-b

	fmt.Printf("x: %d, y: %d, x+y: %d\n", x, y, x+y)

}

func partialSum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}

	c <- sum
}
