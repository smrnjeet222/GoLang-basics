package channels

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

// channels and gorutines
func Channel() {
	ch := make(chan int, 50) //buffer 50
	wg.Add(2)
	go func(ch <-chan int) { //recieve only channel
		// i := <-ch
		// fmt.Println(i)
		// i = <-ch
		// fmt.Println(i)

		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) { //send only channel
		ch <- 42
		ch <- 27
		close(ch) //removes deadlock
		wg.Done()
	}(ch)

	wg.Wait()
}
