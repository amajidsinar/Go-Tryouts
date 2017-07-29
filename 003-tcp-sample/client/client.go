package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalf("%s\r\n", err)
	}
	defer conn.Close()

	d := "I dialed you"
	fmt.Fprintf(conn, "%s\r\n", d)
}
