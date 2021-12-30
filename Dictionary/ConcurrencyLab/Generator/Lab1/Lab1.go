package main

import "strings"

func main() {
	data := []string{"Ace destroy everything around his area",
		"King oredering 5 of powerhouse to keep place very safety and warm",
		"Queen keep making the server peacefull and get well aslong as the world coming",
		"Jack surprising all othe powerhouses",
		"Joker the main brain  and power is with him"}
	histogram := make(map[string]int)
	words := words(data)
	for word := range words {
		histogram[word]++
	}
}
func words(data []string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out) //close channel after returning the function
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word = strings.ToLower(word)
				out <- word
			}
		}
	}()
	return out
}
