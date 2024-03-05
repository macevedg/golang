package main

import (
	"fmt"
)

//given a tree, get the average value at each level
/*
input:

     4
	/ \
	7  9
   / \  \
  10  2  6
       \
	   6
	  /
	  2

output: [4,8,6,6,2]
*/

// breath first search

func main() {
	q := AdvancedQueue{}

	nextLevel := AdvancedQueue{}
	numberOfLevels := 0

	q.add(inicialize())

	avgPerLevel := make([]int64, 1)
	avgPerLevel[numberOfLevels] = 0
	avgPerLevelTotal := 0

	for !q.isEmpty() {
		node := q.pop().(*Node)

		if node == nil {
			if q.isEmpty() {
				q = nextLevel
				nextLevel = AdvancedQueue{}
				if avgPerLevel[numberOfLevels] != 0 {
					avgPerLevel[numberOfLevels] = avgPerLevel[numberOfLevels] / int64(avgPerLevelTotal)
				}
				fmt.Println()

				if !q.isEmpty() {
					numberOfLevels++
					avgPerLevelTotal = 0
					avgPerLevel = append(avgPerLevel, 0)
				}
			}

			continue
		}

		fmt.Printf("%d ", node.data)
		avgPerLevel[numberOfLevels] = (avgPerLevel[numberOfLevels] + node.data)
		avgPerLevelTotal++

		if node.right != nil {
			nextLevel.add(node.right)
		}

		if node.left != nil {
			nextLevel.add(node.left)
		}

		if q.isEmpty() {
			q = nextLevel
			nextLevel = AdvancedQueue{}
			if avgPerLevel[numberOfLevels] != 0 {
				avgPerLevel[numberOfLevels] = avgPerLevel[numberOfLevels] / int64(avgPerLevelTotal)
			}

			fmt.Println()

			if !q.isEmpty() {
				numberOfLevels++
				avgPerLevelTotal = 0
				avgPerLevel = append(avgPerLevel, 0)
			}
		}
	}

	fmt.Printf("Tree has %d levels", numberOfLevels)
	fmt.Printf("output: %v\n", avgPerLevel)

}

type Node struct {
	data int64

	left  *Node
	right *Node
}

type AdvancedNode struct {
	data float64

	left  *Node
	right *Node
}

func inicialize() *Node {
	root := &Node{data: 4}

	root.left = &Node{data: 7}
	root.right = &Node{data: 9}

	left := root.left

	left.left = &Node{data: 10}
	left.right = &Node{data: 2}

	right := root.right
	right.right = &Node{data: 6}

	left = left.right

	left.right = &Node{data: 6}

	left.right.left = &Node{data: 2}

	return root
}

// simple queue
type Queue struct {
	data []int64
}

func (q *Queue) add(i ...int64) {
	q.data = append(q.data, i...)
}

func (q *Queue) pop() int64 {
	if q.isEmpty() {
		return -1
	}
	element := q.data[0]
	q.data = q.data[1:]

	return element
}

func (q *Queue) isEmpty() bool {
	return len(q.data) == 0
}

// advanced queue
type AdvancedQueue struct {
	data []interface{}
}

func (q *AdvancedQueue) add(i ...interface{}) {
	q.data = append(q.data, i...)
}

func (q *AdvancedQueue) pop() interface{} {
	if q.isEmpty() {
		return nil
	}
	element := q.data[0]
	q.data = q.data[1:]

	return element
}

func (q *AdvancedQueue) isEmpty() bool {
	return len(q.data) == 0
}
