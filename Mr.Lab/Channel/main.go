package main

import (
	"bytes"
	"fmt"
	"sync"
)

/*
func main() {
	chanowner := func() <-chan int {
		results := make(chan int, 5) // 1. Prevent other go routine access it
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}
	consumer := func(results <-chan int) { // 3
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}
	results := chanowner() // 2
	consumer(results)
} // simple channel
*/
/*
1. 	Here we instantiate the channel within the lexical scope of the chanowner function.
	This limitd the scope of the write aspect of the results chanel the closure defined below it.
	In other words, it connes the write aspect of this channel to prevent other goroutines from writing to it.

2. 	Here we receive the read aspect of the channel and we're able to pass it into the consumer, which can do nothing but read from it.
	Once again this confines the main goroutine to read-only view of the channel.

3.	Here we receive a read-only copy of an int channel. By declaring that the only usage we require is read access, we confine usage of
	the channel within the consume function to only reads.
*/
func main() {
	var wg sync.WaitGroup
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}
	wg.Add(2)
	data := []byte("GolangGo")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])
}
