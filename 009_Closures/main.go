package main

import "fmt"

func intSec() func() int {
	var i int
	return func() int {
		i += 2
		return i
	}
}

func main() {
	nextInt := intSec()
	for i, _ := range make([]int, 25) {
		fmt.Println(nextInt(), i)
	}
}
