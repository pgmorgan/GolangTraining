package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()
	go func() {
		c <- 43
	}()

	fmt.Println(<-c)
	fmt.Println(<-c)
}
