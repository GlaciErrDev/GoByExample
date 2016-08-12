package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("1)Received message -", msg)
	default:
		fmt.Println("No message received")
	}

	msg := "!!!_FooBarBaz_!!!"
	select {
	case messages <- msg:
		fmt.Println("Sent message -", msg)
	default:
		fmt.Println("No message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("2)Received message -", msg)
	case sig := <-signals:
		fmt.Println("Received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
