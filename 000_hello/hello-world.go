package main

import "fmt"

func main() {
	h := "Hello World"
	// newline is denoted by \r\n in windows
	// and denoted by \n in mac
	fmt.Printf("%s\r\n", h)
}
