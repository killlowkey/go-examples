package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("debugger")
		time.Sleep(5 * time.Second)
	}
}
