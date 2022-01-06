package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan int)

	go func() {
		ch1 <- "derek"
		ch2 <- 111
	}()

	go func() {
		content := <-ch1
		fmt.Println(" take out the data ：", content) // take out the data ： derek
		ch2 <- 222
	}()

	a := <-ch2
	b := <-ch2
	fmt.Println(a, b) //111 222
	fmt.Println(" end of the program ")
}
