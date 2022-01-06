package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			fmt.Println(" the first ", j, " time to perform ")
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(" end of the program ")
}
