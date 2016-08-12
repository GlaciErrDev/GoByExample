package main

import (
	"fmt"
)

func main() {
	queue := make(chan string, 2)
	queue <- "Foo"
	queue <- "Bar"
	close(queue)

	for el := range queue {
		fmt.Println(el)
	}
}
