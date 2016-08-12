package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("!!!FooBarBaz!!!")

	os.Exit(3)
}
