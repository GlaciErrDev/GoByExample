package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- "Foo!"
	}()

	select {
	case res := <-ch1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout ch1")
	}

	ch2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "Bar!!!"
	}()

	select {
	case res := <-ch2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout ch2")
	}

}
