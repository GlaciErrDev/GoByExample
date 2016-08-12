package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func main() {
	signals := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
	    sig := <-signals
		fmt.Println()
		fmt.Println(sig)
		done<- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
