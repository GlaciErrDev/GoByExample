package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for i, n := range nums {
		fmt.Println("i ->", i)
		fmt.Println("num ->", n)
		sum += n
	}
	fmt.Println("Sum ->", sum)

	kvs := map[string]string{"foo": "bar", "baz": "xyz"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "Go" {
		fmt.Println(i, c)
	}
}
