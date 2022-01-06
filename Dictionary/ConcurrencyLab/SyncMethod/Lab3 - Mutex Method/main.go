package main

import (
	"fmt"
	"sync"
)

func main() {
	var rwm sync.RWMutex
	var wg sync.WaitGroup
	wg.Add(10)
	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		go func(j int) {
			rwm.Lock()
			m[j] = j
			fmt.Println(m)
			rwm.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(" end of the program ")
}
