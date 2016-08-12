package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Foo", age: 42})
	fmt.Println(person{name: "BobdoD"})
	fmt.Println("Struct ptr:", &person{name: "Ann", age: 30})

	s := person{name: "Tom", age: 25}
	fmt.Println(s.age, s.name)

	sptr := &s
	fmt.Println(sptr.name, sptr.age)

	sptr.age = 30
	fmt.Println(*sptr)
}
