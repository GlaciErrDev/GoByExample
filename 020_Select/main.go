package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "FOO"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "BAR"
	}()

	go func() {
		time.Sleep(time.Second * 3)
		ch3 <- "BAZ"
	}()

	t1 := time.Now()
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("received", msg1)
		case msg2 := <-ch2:
			fmt.Println("received", msg2)
		case msg3 := <-ch3:
			fmt.Println("received", msg3)
		}
	}
	fmt.Println("real ->", time.Since(t1))
}
