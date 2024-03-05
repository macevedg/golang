/*
package concurrency

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu = sync.Mutex{}
)

type myNum struct {
	data int
}

func mainold() {

	initialSum := myNum{
		data: 20,
	}
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go changeit(&initialSum, i, wg)
	}

	wg.Wait()
	fmt.Println("total sum: ", initialSum)
}

func changeit(sum *myNum, newValue int, wg *sync.WaitGroup) {
	defer wg.Done()

	//mu.Lock()
	sum.data += newValue
	//mu.Unlock()
}

func channelstest() {
	counter := 10
	now := time.Now()
	ch := make(chan string)
	wg := &sync.WaitGroup{}

	for i := 0; i < counter; i++ {
		wg.Add(1)
		go printit(wg, i+1, ch)
	}
	//wg.Wait()
	for i := 0; i < counter; i++ {
		val, ok := <-ch

		if !ok {
			fmt.Println("i break out of this thing")
			break
		}
		time.Sleep(time.Second)
		fmt.Printf("[%v] - %s <%t>\n", time.Now(), val, ok)
	}

	wg.Wait()
	close(ch)
	fmt.Println("Elapsed time: ", time.Since(now))
}

func printit(wg *sync.WaitGroup, counter int, ch chan<- string) {
	defer endit(wg, ch)

	time.Sleep(time.Second * time.Duration(counter))

	ch <- fmt.Sprintf("%v", counter)
}

func endit(wg *sync.WaitGroup, ch chan<- string) {
	wg.Done()
}
*/