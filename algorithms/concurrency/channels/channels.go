package channels

import (
	"fmt"
	"sync"
)

func Execute(ch chan string) {
	wg := sync.WaitGroup{}
	//mu := sync.Mutex{}
	for i := 0; i < 10; i++ {
		param := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- fmt.Sprintf("number %d", param)
			fmt.Printf("Execute ch:%v\n", ch)
		}()
	}

	wg.Wait()
	ch <- "quit"
}
