package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

//➜ echo 'hello'   > /tmp/lines
//➜ echo 'filter' >> /tmp/lines
//➜ cat /tmp/lines | go run 051_LineFilters/main.go
//HELLO
//FILTER


func main() {
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		upline := strings.ToUpper(scan.Text())
		fmt.Println(upline)
	}

	if err := scan.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
