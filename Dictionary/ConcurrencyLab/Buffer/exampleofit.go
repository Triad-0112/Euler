package main

import "fmt"

func main() {
	buff := 1000
	ch := make(chan int, buff) // creating 3 buffer which you can analogy of it as 3 container who will used for communicating in channel, let say a frequency of radio
	for i := 0; i < buff; i++ {
		ch <- i
	}
	for i := 0; i < buff; i++ {
		fmt.Println(<-ch)
	}
}
