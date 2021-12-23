package main

import "fmt"

func main() {
	a := 0
	b := 1
	total := 0
	const max = 4000000
	for b < max {
		mem := a + b
		if b%2 == 0 {
			total += b
		}
		a = b
		b = mem
	}
	fmt.Printf("Total nya adalah %d", total)
}
