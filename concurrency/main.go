package main

import (
	"fmt"
	"time"
)

func count(thing string) {
	for i := 0; i < 5; i++ {

		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

//wait group
/*
func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		count("sheep")
	}()

	go func() {
		defer wg.Done()

		count("leopard")
	}()

	wg.Wait()

	//fmt.Scanln()
}
*/

// channels
func main() {

	c := make(chan string)

	go countch("sheep", c)

	for msg := range c {

		fmt.Println(msg)
	}
}

func countch(thing string, c chan string) {
	for i := 0; i < 5; i++ {

		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
