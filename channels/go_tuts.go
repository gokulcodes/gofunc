package main

import (
	"fmt"
	"time"
)

func valueParser(c chan string) {
	c <- "myval"
	fmt.Println("Received", c)
	// time.Sleep(2 * time.Second)
	fmt.Println("done")
}

func main() {
	vals := make(chan string, 2)
	// fmt.Println("Execution of go-Routines")
	go valueParser(vals)
	go valueParser(vals)
	val, val1 := <-vals, <-vals
	fmt.Println(val, val1)
	time.Sleep(2 * time.Second)
}
