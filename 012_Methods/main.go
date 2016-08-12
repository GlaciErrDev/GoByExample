package main

import "fmt"

type rectangle struct {
	width, height int
}

func (r rectangle) area() int {
	return r.height * r.width
}

func (r rectangle) perimetr() int {
	return 2*r.height + 2*r.width
}

func main() {
	r := rectangle{width: 10, height: 5}
	fmt.Println("Area ->", r.area())
	fmt.Println("Perimeter ->", r.perimetr())

	rptr := &r
	fmt.Println("Area ->", rptr.area())
	fmt.Println("Perimeter ->", rptr.perimetr())
}
