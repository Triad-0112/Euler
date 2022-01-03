package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan string, results chan<- string) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "started job", j)
		results <- j + " job done"
	}
}
func main() {
	jobsList := []string{"job1", "job2", "job3", "job4", "job5",
		"job6", "job7", "job8", "job9", "job10",
		"job11", "job12", "job13", "job14", "job"}
	numJobs := len(jobsList)
	jobs := make(chan string, numJobs)
	results := make(chan string, numJobs)
	for w := 1; w <= 6; w++ {
		go worker(w, jobs, results)
	}
	for _, job := range jobsList {
		jobs <- job
	}
	close(jobs)
	for a := 1; a <= numJobs; a++ {
		fmt.Println(<-results)
	}
}
