package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func BubbleSortVanilla(intList []int) {
	for i := 0; i < len(intList)-1; i += 1 {
		if intList[i] > intList[i+1] {
			intList[i], intList[i+1] = intList[i+1], intList[i]
		}
	}
}

func BubbleSortOdd(intList []int, wg *sync.WaitGroup, c chan []int) {
	for i := 1; i < len(intList)-2; i += 2 {
		if intList[i] > intList[i+1] {
			intList[i], intList[i+1] = intList[i+1], intList[i]
		}
	}
	wg.Done()
}

func BubbleSortEven(intList []int, wg *sync.WaitGroup, c chan []int) {
	for i := 0; i < len(intList)-1; i += 2 {
		if intList[i] > intList[i+1] {
			intList[i], intList[i+1] = intList[i+1], intList[i]
		}
	}
	wg.Done()
}

func ConcurrentBubbleSort(intList []int, wg *sync.WaitGroup) {
	for i := 0; i < len(intList)-1; i++ {
		if intList[i] > intList[i+1] {
			intList[i], intList[i+1] = intList[i+1], intList[i]
		}
	}
	wg.Done()
}

func main() {
	// defer profile.Start(profile.MemProfile).Stop()
	rand.Seed(time.Now().Unix())
	intList := rand.Perm(100000)
	fmt.Println("Read a sequence of", len(intList), "elements")

	//c := make(chan []int, len(intList))
	var wg sync.WaitGroup

	start := time.Now()
	for j := 0; j < len(intList)-1; j++ {
		// BubbleSortVanilla(intList) // takes roughly 15s

		// wg.Add(2)
		// go BubbleSortOdd(intList, &wg, c)  // takes roughly 7s
		// go BubbleSortEven(intList, &wg, c)
		wg.Add(1) // takes roughly 1.4s
		go ConcurrentBubbleSort(intList, &wg)
	}
	wg.Wait()
	elapsed := time.Since(start)

	// Print the sorted integers
	fmt.Println("Sorted List: ", len(intList), "in", elapsed)
	fmt.Println(intList)
}
