package main

import (
	"fmt"
//	"time"
)

func routine(c chan string){
	c <- "Value Passed to channel"	
	fmt.Println(c)
}

func main(){
	val := make(chan string, 2)
	fmt.Println("Go channing......")	
	go routine(val)
	go routine(val)
//	time.Sleep(4 * time.Second)
	value := <- val
	fmt.Println("Received\n", value)
	fmt.Println(<-val)
}
