// package main

// import (
// 	"fmt"
// 	"net"
// )

// func strlen(s string, c chan int) {
// 	fmt.Println(c)
// 	c <- len(s)
// }

// func main() {
// 	c := make(chan int)
// 	go strlen("name", c)
// 	go strlen("id", c)
// 	x, y := <-c, <-c
// 	fmt.Println(x, y)
// 	con, err := net.Dial("tcp", "scanme.nmap.org:80")
// 	if err == nil {
// 		fmt.Println("Successfull", con)
// 	} else {
// 		fmt.Println("Unsuccessfull", con)
// 	}
// }
