package main

import "fmt"

//check below, how channel work by take a look at their length for each value that taken from the container
func main() {
	ch := make(chan int, 10)
	makeEvemNums(5, ch)
	fmt.Printf("check length and capacity of channel: %d, capacity: %d\n", len(ch), cap(ch))
	fmt.Println(<-ch)
	fmt.Printf("check length and capacity of channel: %d, capacity: %d\n", len(ch), cap(ch))
	fmt.Println(<-ch)
	fmt.Printf("check length and capacity of channel: %d, capacity: %d\n", len(ch), cap(ch))
	fmt.Println(<-ch)
	fmt.Printf("check length and capacity of channel: %d, capacity: %d\n", len(ch), cap(ch))
	fmt.Println(<-ch)
	fmt.Printf("check length and capacity of channel: %d, capacity: %d\n", len(ch), cap(ch))
	fmt.Println(<-ch)
	fmt.Printf("check length and capacity of channel: %d, capacity: %d\n", len(ch), cap(ch))
}
func makeEvemNums(count int, in chan<- int) {
	for i := 0; i < count; i++ {
		in <- 2 * i
	}
}
