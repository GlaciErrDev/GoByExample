package main

import "fmt"

func sum(nums ...int) int {
	var res int
	for _, n := range nums {
		res += n
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sum(nums...))
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(5, 5, 5, 5, 5))
}
