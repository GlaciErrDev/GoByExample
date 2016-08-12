package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	name          string
	width, height float64
}

type circle struct {
	name   string
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return 2*r.height + 2*r.width
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("Figure ->", g)
	fmt.Println("Area ->", g.area())
	fmt.Println("Perimeter ->", g.perim())
}

func main() {
	r := rect{width: 10, height: 5, name: "rectangle"}
	c := circle{radius: 15, name: "circle"}

	measure(r)
	measure(c)
}
