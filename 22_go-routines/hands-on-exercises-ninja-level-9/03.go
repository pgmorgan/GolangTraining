package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var mutex sync.Mutex
	var wg sync.WaitGroup

	value := 0
	myRange := 99
	wg.Add(myRange)

	for i := 0; i < myRange; i++ {
		go func() {
			fmt.Println(runtime.NumGoroutine())
			mutex.Lock()
			value++
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final value:", value)
}
