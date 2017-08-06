package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	//In client server interaction, the server always waits for request from client
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		//Accept all incoming connection
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		//Handle the connection in a new goroutine
		//The loop then returns to accepting
		//so that handle may be served concurrently
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	request(conn)

	respond(conn)
}

func request(conn net.Conn) {
	//Scan whole page
	scanner := bufio.NewScanner(conn)
	//Scan for each token(line)
	for scanner.Scan() {
		//Pass them into each line
		ln := scanner.Text()
		fmt.Println(ln)
	}
}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}
