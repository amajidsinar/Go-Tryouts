package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//Connect to port 8080
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalf("%s\r\n", err)
	}
	defer conn.Close()

	//Pass d into conn
	d := "I dialed you"
	fmt.Fprintf(conn, "%s\r\n", d)
}
