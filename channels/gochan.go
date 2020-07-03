package main


import (
	"fmt"
	"net/http"
//	"sync"
	"log"
	"runtime"
	"golang.org/x/crypto/acme"
)

func routine(c chan string){
	c <- "https://www.netflix.com"
	fmt.Printf("%+v", c)
//	defer wg.Done()
	fmt.Printf("Done")
}


func main(){
//	var wg sync.WaitGroup
	com :=  make(chan string, 2)
	fmt.Println("Initialized...")
//	wg.Add(1)
	go routine(com)
	routine(com)
	x := <-com
//	fmt.Println(com, x)
	cont, err := http.Get(x)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cont)
	fmt.Println(runtime.GOMAXPROCS(2))
//	wg.Wait()
}
