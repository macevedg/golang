package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("Hello world!")

	// explicit
	//var human Human

	human1 := User{
		Human: Human{
			name: "John Doe",
		},
		age: 11,
	}

	human2 := User{
		Human: Human{
			name: "Jane Doe",
		},
		age: 10,
	}

	fmt.Printf("available CPUS: %v\n", runtime.NumCPU())

	now := time.Now()
	wg := &sync.WaitGroup{}

	wg.Add(2)
	var val1, val2 int
	go func() {
		val1 = processUser(human1, wg)
	}()

	go func() {
		val2 = processUser(human2, wg)
	}()

	wg.Wait() //

	fmt.Printf("elapsed time: %v \n", time.Since(now))
	fmt.Printf("result: %d \n", (val1 + val2))
}

func processUser(user User, wg *sync.WaitGroup) int {

	for i := 0; i < 10; i++ {

		fmt.Println(user.name)
	}
	time.Sleep(2 * time.Second)

	wg.Done()

	return user.age
}

type Human struct {
	name string
}

type User struct {
	Human
	age int
}
