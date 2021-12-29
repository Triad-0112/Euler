package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string, 3)

	messages <- "one"
	messages <- "two"
	messages <- "three"

	go func(m *chan string) {
		fmt.Println("Entering the goroutine...")
		for {
			fmt.Println(<-*m)
		}
	}(&messages)
	time.Sleep(5 * time.Second)
	fmt.Println("Done!")
}

// https://go.dev/tour/concurrency/5
