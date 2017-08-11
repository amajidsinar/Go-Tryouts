package main

import (
	"fmt"
	"net/http"
)

type Hotdog int

func (h Hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hotdog")
}

type Burger int

func (b Burger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "burger")
}

func main() {
	var h Hotdog
	var b Burger

	server := http.Server{
		Addr: ":8080",
	}

	http.Handle("/hotdog", h)
	http.Handle("/burger", b)

	server.ListenAndServe()
}
