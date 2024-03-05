package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type person struct {
	ID      int
	Level   int
	Profile string
}

func main() {
	function1()

	time.Sleep(10000 * time.Millisecond)
}

func function1() {
	go gofunction()

	fmt.Println("finished function1")
}

func gofunction() {
	for {
		fmt.Println("goroutine executing...")
		time.Sleep(1000 * time.Millisecond)
	}
}

func testString(selects ...string) {

	fmt.Printf("length: %d\n", len(selects))
}

func mainoldd() {

	mat := make(map[int]map[int]string, 0)

	list := []person{
		person{
			ID:      1,
			Level:   1,
			Profile: "profile11",
		},
		person{
			ID:      1,
			Level:   2,
			Profile: "profile12",
		},
		person{
			ID:      2,
			Level:   2,
			Profile: "profile22",
		},
		person{
			ID:      2,
			Level:   3,
			Profile: "profile23",
		},
		person{
			ID:      2,
			Level:   4,
			Profile: "profile24",
		},
		person{
			ID:      3,
			Level:   1,
			Profile: "profile31",
		},
	}

	for _, p := range list {

		if mat[p.ID] == nil {
			mat[p.ID] = make(map[int]string, 0)
		}

		mat[p.ID][p.Level] = p.Profile
	}

	b, _ := json.Marshal(&mat)

	fmt.Println("JSON marshaling: %s", string(b))

	for i, m := range mat {
		for j, m2 := range m {
			fmt.Printf("ID:%d, Level: %d = %s \n", i, j, m2)
		}
	}

	fmt.Println("yes: %v\n", mat)
}

func mainOLD() {

	l := append([]int{}, []int{8, 2, 3, 4, 5, 6, 7, 8, 9}...)

	size := len(l)
	for i, n := range l {

		right := i + n

		if right >= size {
			right = size
		}

		mid := l[i:right]

		fmt.Printf("i=%d, right=%d, array: %v\n", i, right, mid)
	}

}

func getMinAndMAx(P int, Q int, A []int) (int, int) {

	//if P == Q {
	//    return (A[P], A[Q])
	//}
	//check := make(map[int]int, 0)

	minn := 100000000000
	maxx := 0

	for i := P; i < Q; i++ {
		if minn > A[i] {
			minn = A[i]
		}

		if maxx < A[i] {
			maxx = A[i]
		}
	}

	return minn, maxx
}
