package main

import (
	"fmt"
	"os"
)

func createFile(s string) *os.File {
	fmt.Println("Creating file.")
	f, err := os.Create(s)
	if err != nil {
		panic(err)
	}
	return f
}

func closeFile(f *os.File) {
	fmt.Println("Closing file.")
	f.Close()
}

func writeTextToFile(f *os.File, s string) {
	fmt.Println("Writing to file ->", s)
	fmt.Fprintln(f, s)
}

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeTextToFile(f, "FooBarBaz!!!")
}
