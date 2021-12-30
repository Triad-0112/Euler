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
	done := make(chan struct{})
	go func() {
		defer close(done)
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word = strings.ToLower(word)
				histogram[word]++
			}
		}
	}()
	<-done
	for k, v := range histogram {
		fmt.Printf("%s\t(%d)\n", k, v)
	}
}
