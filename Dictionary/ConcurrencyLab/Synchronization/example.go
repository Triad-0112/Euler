package main

import (
	"fmt"
	"strings"
)

// warning: this example can cause some race condition
func main() {
	data := []string{"Ace destroy everything around his area",
		"King oredering 5 of powerhouse to keep place very safety and warm",
		"Queen keep making the server peacefull and get well aslong as the world coming",
		"Jack surprising all othe powerhouses",
		"Joker the main brain  and power is with him"}
	histogram := make(map[string]int)
	done := make(chan bool)
	go func() {
		for _, line := range data {
			words := strings.Split(line, " ") // this will run concurrently, so which stirng that will going to be processed first and who is the second is random
			for _, word := range words {
				word = strings.ToLower(word)
				histogram[word]++
			}
		}
		done <- true
	}()
	if <-done {
		for k, v := range histogram {
			fmt.Printf("%s\t(%d)\n", k, v)
		}
	}

}

/*
If you see data above, you will see how go concurrency work. Its completely random which will be run in compiler, which was not.
Go Routines is processing data together but with different execution thread that going to be used for them to store the value in
memory address.
*/
