// package main

// import (
// 	"fmt"
// 	"time"
// )

// func wrriter(c chan string) {
// 	c <- "something"
// 	time.Sleep(10 * time.Second)
// }

// func main() {
// 	ch := make(chan string)
// 	fmt.Println("Channel Implementation")
// 	go wrriter(ch)
// 	fmt.Println(<-ch)
// 	fmt.Println("Go Routine 1")
// 	go wrriter(ch)
// 	fmt.Println(<-ch)
// 	fmt.Println("Go Routine 2")
// }
