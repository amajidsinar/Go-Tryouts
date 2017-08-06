package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	//Listen on port 8080
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicf("%s\r\n", err)
	}
	defer li.Close()

	for {
		//Accept all oncoming connection to port 8080
		conn, err := li.Accept()
		if err != nil {
			log.Printf("%s\r\n", err)
		}
		defer conn.Close()

		//Handle the connection in a new goroutine
		//The loop then returns to accepting, so that
		//multiple connection may be served concurrently
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	//Scan(read) what is written in conn
	scanner := bufio.NewScanner(conn)
	//Scan for every token
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Printf("%s\r\n", ln)
	}
	defer conn.Close()
}
