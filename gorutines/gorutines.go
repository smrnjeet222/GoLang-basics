package gorutines

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

// Gorutines is.....
func Gorutine() {
	runtime.GOMAXPROCS(10)
	var msg = "Hello Go!"
	wg.Add(1)
	go func(msg string) { // dont use msg of outer scope / pass in func
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "Goodbye"
	// time.Sleep(100 * time.Millisecond)

	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	// m.RLock()
	fmt.Println("Hello", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	// m.Lock()
	counter++
	m.Unlock()
	wg.Done()
}
