package main

import "fmt"

const RANGE int = 100

func main() {
	c := make(chan int)

	go foo(c)
	// go func() {
	// 	for i := 0; i < RANGE; i++ {
	// 		c <- i
	// 	}
	// 	close(c)
	// }()
	// bar(c)
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit.")
}

func foo(c chan<- int) {
	for i := 0; i < RANGE; i++ {
		c <- i
	}
	close(c)
}

// func bar(c <-chan int) {
// 	for v := range c {
// 		fmt.Println(v)
// 	}
// }
