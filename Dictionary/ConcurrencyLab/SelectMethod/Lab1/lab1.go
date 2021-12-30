package main

import (
	"fmt"
	"strings"
)

func main() {
	data := []string{"Ace destroy everything around his area",
		"King oredering 5 of powerhouse to keep place very safety and warm",
		"Queen keep making the server peacefull and get well aslong as the world coming",
		"Jack surprising all othe powerhouses",
		"Joker the main brain  and power is with him"}
	histogram := make(map[string]int)
	stopCh := make(chan struct{}) // used to signal stop
	words := words(stopCh, data)  // returns handle to channel
	for word := range words {
		if histogram["the"] == 3 {
			close(stopCh)
		}
		histogram[word]++
	}
	fmt.Println(data)
	fmt.Println(histogram)
	fmt.Println(<-words) // return nothing...
}

func words(stopCh chan struct{}, data []string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out) // closes channel upon fn return
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word = strings.ToLower(word)
				select {
				case out <- word:
				case <-stopCh: // succeeds first when close
					return
				}
			}
		}
	}()
	return out
}
