package main

import (
	"os/exec"
	"fmt"
	"io/ioutil"
)

func main() {
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("date >\n", string(dateOut))

	grepCmd := exec.Command("grep", "grep")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
    grepCmd.Wait()
	fmt.Println("grep >\n", string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -lah")
	lsOut, err := lsCmd.Output()
	if err != nil {
        panic(err)
    }
    fmt.Println("ls -lah >\n", string(lsOut))
}
