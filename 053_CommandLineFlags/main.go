package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("bool", false, "a bool")

	var strVar string
	flag.StringVar(&strVar, "strVar", "bar", "a string var")

	flag.Parse()

	//fmt.Print("ptr: ", &wordPtr, " - ")
	fmt.Println("word:", *wordPtr)
    //fmt.Print("ptr: ", &numbPtr, " - ")
    fmt.Println("numb:", *numbPtr)
    //fmt.Print("ptr: ", &boolPtr, " - ")
    fmt.Println("bool:", *boolPtr)

	fmt.Println("svar:", strVar)

	fmt.Println("tail:", flag.Args())
}
