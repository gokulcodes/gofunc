package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			con, err := net.Dial("tcp", address)
			fmt.Println(con)
			if err != nil {
				// fmt.Println(err)
				return
			}
			con.Close()
			fmt.Println("open:%d", j)
		}(i)
	}
	wg.Wait()
}
