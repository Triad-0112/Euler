package main

import "fmt"

//Inproper way!!
func main() {
	ch := make(chan int)
	// ch <- 12 // blocks, this is was unbuffered channel which will be waiting for value been taken from channel
	go func() { ch <- 12 }()
	fmt.Println(<-ch)
}
