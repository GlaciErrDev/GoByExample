package main

import (
	"fmt"
	//"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("Worker id -", id, "processing job -", j)
		//time.Sleep(time.Second)
		result <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for wId := 1; wId <= 3; wId++ {
		go worker(wId, jobs, results)
	}

	last := 9

	for j := 1; j <= last; j++ {
		jobs <- j
	}

	close(jobs)

	for r := 1; r <= last; r++ {
		<-results
	}
}
