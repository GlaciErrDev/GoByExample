package main

import "fmt"

func vals() (int, int) {
	return 5, 10
}

func main() {
	a, b := vals()
	fmt.Println(a, b)
	_, y := vals()
	fmt.Println(y)
}
