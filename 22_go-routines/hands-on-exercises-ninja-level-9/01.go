package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Num CPU Cores:", runtime.NumCPU())
	fmt.Println("Num Go Routine: ", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hello from 1.")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from 2.")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("About to exit.")
}
