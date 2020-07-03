package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func valueParser(c chan string) {
	c <- "myval"
	fmt.Println("Received", c)
	// time.Sleep(2 * time.Second)
	fmt.Println("done")
	defer wg.Done()
}

func main() {
	vals := make(chan string)
	// fmt.Println("Execution of go-Routines")
	wg.Add(1)
	go valueParser(vals)
	wg.Add(1)
	go valueParser(vals)
	val, val1 := <-vals, <-vals
	fmt.Println(val, val1)

	wg.Wait()
}
