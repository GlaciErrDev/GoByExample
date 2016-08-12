package main

import (
	"errors"
	"fmt"
)

func foo(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42!!!")
	}
	return arg * arg, nil
}

type argError struct {
	arg  int
	prob string
}

func (e argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func bar(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "Can't work with it"}
	}
	return arg + arg, nil
}

func main() {
	for _, i := range []int{5, 42} {
		if r, e := foo(i); e != nil {
			fmt.Println("'foo' failed ->", e)
		} else {
			fmt.Println("'foo' result ->", r)
		}
	}

	for _, i := range []int{5, 42} {
		if r, e := bar(i); e != nil {
			fmt.Println("'foo' failed ->", e)
		} else {
			fmt.Println("'foo' result ->", r)
		}
	}

	_, e := bar(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ok)
		fmt.Println(ae)
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
