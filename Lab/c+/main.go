package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func start() {
	dataLen := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range dataLen {
		go func(temp int) {
			fmt.Println(temp)
		}(i)
	}
}

func main() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		start()
		<-sigs
		done <- true
	}()

	<-done
}
