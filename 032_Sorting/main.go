package main

import (
	"fmt"
	"sort"
)

func main() {
	strArr := []string{"c", "a", "b"}
	fmt.Println(strArr)
	sort.Strings(strArr)
	fmt.Println(strArr)

	intarr := []int{5, 3, 7, 5, 9, 3, 6, 2}
	fmt.Println(intarr)
	sort.Ints(intarr)
	fmt.Println(intarr)

	fmt.Println("Int sorted:", sort.IntsAreSorted(intarr))
}
