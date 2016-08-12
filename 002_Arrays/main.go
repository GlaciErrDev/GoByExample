package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var a [5]int
	fmt.Println("empty array:", a)

	a[4] = 100
	fmt.Println("set a[4]:", a)
	fmt.Println("get a[4]:", a[4])
	fmt.Println("len a:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("declare b:", b)

	var doubleArray [3][5]int
	for i := 0; i < len(doubleArray); i++ {
		for j := 0; j < len(doubleArray[i]); j++ {
			doubleArray[i][j] = i + j
		}
	}
	fmt.Println("double array:", doubleArray)
}
