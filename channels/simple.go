package main

import (
	"fmt"
	"runtime"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
}

func outputText(j *Job) {
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(j.text)
		j.i++
	}
	// goGroup.Done()
}
func main() {
	// goGroup := new(sync.WaitGroup)
	hello := new(Job)
	world := new(Job)
	hello.text = "hello"
	hello.i = 0
	hello.max = 3
	world.text = "world"
	world.i = 0
	world.max = 5
	// goGroup.Add(1)
	go outputText(hello)
	go outputText(world)
	runtime.Gosched()
	fmt.Printf("%+v\n", runtime.GOMAXPROCS(2))
	fmt.Println(time.Now().UnixNano())
	val := []byte(" ")
	fmt.Println(string(val))
	// goGroup.Wait()
}
