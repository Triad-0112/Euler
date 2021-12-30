package main

import (
	"fmt"
)

func foo(c chan int) {

	for i := 0; i <= cap(c); i++ {
		// how to close the channel after 100 integers are written?
		c <- i + 1
		if i == 49 {
			close(c)
			break
		}
		c <- i + 2

	}
}

func main() {
	c := make(chan int, 200)
	c <- 0
	go foo(c)
	for i := 0; i < cap(c); i++ {
		fmt.Println(<-c)
	}
}
