package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	//done := make(chan bool)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
		//done <- true
	}()

	stop := timer2.Stop()
	if stop {
		fmt.Println("Timer 2 stopped")
	}
	//<-done
}
