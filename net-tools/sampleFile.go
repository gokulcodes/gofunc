package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}
type Dog struct{}
type Friend interface {
	sayHello()
}

func slicer() {
	var s = make([]string, 1)
	var m = make(map[string]string)
	s = append(s, "something")
	m["indexer"] = "sample"
	fmt.Println(s, m)
}
func points() {
	var count = int(42)
	ptr := &count
	fmt.Println(*ptr)
	*ptr = 10000
	fmt.Println(count)
}
func (p *Person) structor() {
	fmt.Println("Hello, ", p.Name)
}
func (d *Dog) sayHello() {
	fmt.Println("Dog")
}
func Greet(f Friend) {
	fmt.Println(f)
	f.sayHello()
}

func main() {
	msg := "Hello world in Go"
	for i := 0; i < 2; i++ {
		fmt.Println(msg)
		fmt.Printf("\n")
	}
	slicer()
	points()
	newStruct := new(Person)
	newStruct.Name = "John"
	newStruct.structor()
	// Greet(newStruct)
	dog := new(Dog)
	Greet(dog)
}
