// givin n, return the spiral 2D matrix
// 123
// 894
// 765
// if im going right => going down
// if im going down => going left
// if im going left => going up
// if im going up => going right

// i think i need to intialize the matrix with some default values to check
// on every step you need to check if the position has the default value
// if not you need to doble check the next posible direction
// if there are no more directions avaiable stop

package main

import "fmt"

func main() {
	n := 3

	matrix := [3][3]float64{}

	//initialize the matrix
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = -1
		}
	}

	//loop until there are no more directions avaiable
	baseColumn := []int{1, 0, -1, 0}
	baseRow := []int{0, 1, 0, -1}
	limit := n * n
	currentNumber := 0

	dir := 0
	r, c := 0, 0
	for currentNumber < limit {
		matrix[r][c] = float64(currentNumber)

		r += baseRow[dir]
		c += baseColumn[dir]

		if isInvalid(matrix, r, c) {
			r -= baseRow[dir]
			c -= baseColumn[dir]

			dir = (dir + 1) % 4

			r += baseRow[dir]
			c += baseColumn[dir]
		}
		currentNumber++
	}

	fmt.Printf("final Matrix: %v\n", matrix)
}

func isInvalid(matrix [3][3]float64, r, c int) bool {
	if r < 0 || c < 0 || r >= len(matrix) || c >= len(matrix) {
		return true
	}

	if matrix[r][c] != -1 {
		return true
	}

	return false
}
