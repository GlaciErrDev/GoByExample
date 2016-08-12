package main

import "fmt"

func main() {
	//message := make(chan string, 2)
	message := make(chan string)

	//message <- "foo"
	//message <- "bar"

	go func() { message <- "FOO" }()
	go func() { message <- "BAR" }()
	go func() { message <- "BAR" }()
	go func() { message <- "BAR" }()
	go func() { message <- "FOO" }()
	go func() { message <- "FOO" }()

	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}
