package main

import "fmt"

func main() {
	var data byte = 0x00
	fmt.Printf("%08b ", data)

	// compress type
	data = (data &^ 0x1C) | ((1 << 2) & 0x1C)
	fmt.Printf("%08b ", data)
}
