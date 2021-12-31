package main

import (
	"log"
	"os"
	"strconv"
	"sync"
)

type Worker struct {
	wLog *log.Logger
	Name string
}

func main() {
	totalGoroutines := 500
	done := make(chan bool)
	var wg sync.WaitGroup
	for i := 0; i < totalGoroutines; i++ {
		wg.Add(1)
		myWorker := Worker{}
		myWorker.Name = "Goroutine " + strconv.Itoa(i) + " " + strconv.FormatInt(int64(i), 10) + ""
		myWorker.wLog = log.New(os.Stderr, myWorker.Name, 1)
		go func(w *Worker) {
			w.wLog.Print("Hmm")
			done <- true
		}(&myWorker)
		wg.Done()
	}
	wg.Wait()
	log.Println("...")
	<-done
}
