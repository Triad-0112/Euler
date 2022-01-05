package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int)
	go func() {
		b := <-ch
		fmt.Println(b)
	}()
	ch <- 10
	close(ch)
	time.Sleep(3 * time.Second)

}
