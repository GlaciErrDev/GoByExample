package main

import "os"

func main() {
	panic("WOOOOOOOOOT!!!!")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
