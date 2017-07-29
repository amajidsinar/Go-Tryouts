package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//generate random seed
	//time.Now() is used because time always changing
	//if we put static number the print value stays the same
	rand.Seed(time.Now().Unix())
	r := rand.Intn(3)
	fmt.Printf("%d\r\n", r)
}
