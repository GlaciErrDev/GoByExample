package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second * 2)
	fmt.Println("done")
	done <- true
}

func main() {
	done := make(chan bool)
	go worker(done)

	//<-done
	fmt.Println(<-done)
}
