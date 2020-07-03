package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	c float64
}

func (r rect) area() float64 {
	fmt.Println(r)
	return r.width + r.height
}

func (c circle) area() float64 {
	fmt.Println(c)
	return math.Pi * c.c
}
func (r rect) perim() float64 {
	fmt.Println(r)
	return r.width + r.height
}

func (c circle) perim() float64 {
	fmt.Println(c)
	return math.Pi * c.c
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 50, height: 80}
	c := circle{c: 20}
	measure(r)
	measure(c)
}
