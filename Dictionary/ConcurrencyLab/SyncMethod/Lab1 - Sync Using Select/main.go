package main

import "fmt"

func main() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func(j int) {
			ch <- j
		}(i)
	}
	// use the for loop to keep accepting
	for {
		select {
		case a := <-ch:
			fmt.Println(a)
		default:
		}
	}
}
