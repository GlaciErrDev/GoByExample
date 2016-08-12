package main

import (
	"os/exec"
	"os"
	"syscall"
	"fmt"
)

func main() {
	bin, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-lah"}
	env := os.Environ()

	if err := syscall.Exec(bin, args, env); err != nil {
		panic(err)
	}

	fmt.Println("This line will not execute")
}
