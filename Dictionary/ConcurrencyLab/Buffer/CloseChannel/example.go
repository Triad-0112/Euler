package main

import "fmt"

func main() {
	ch := make(chan int, 4)
	fmt.Printf("check length and capacity of channel: %d, capacity: %d\n", len(ch), cap(ch))
	ch <- 5
	fmt.Printf("check length and capacity of channel: %d, capacity: %d\n", len(ch), cap(ch))
	ch <- 6
	fmt.Printf("check length and capacity of channel: %d, capacity: %d\n", len(ch), cap(ch))
	ch <- 7
	fmt.Printf("check length and capacity of channel: %d, capacity: %d\n", len(ch), cap(ch))
	close(ch)
	//fmt.Printf("check length and capacity of channel: %d, capacity: %d\n", len(ch), cap(ch))
	//ch <- 8 // this will be run panic, becase the channel is closed which was not accepting any values again
	fmt.Printf("check length and capacity of channel: %d, capacity: %d\n", len(ch), cap(ch))
	for i := 0; i < cap(ch); i++ {
		if val, opened := <-ch; opened {
			fmt.Printf("length: %d, capacity: %d\nch: %d\n", len(ch), cap(ch), <-ch)
			fmt.Println(val)
		} else {
			fmt.Println("Channel closed!")
		}
	}

}
