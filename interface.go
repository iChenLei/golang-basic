package main

import "fmt"
import "math"

type geo interface {
	area() float64
	perim() float64
}

type square struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.width * s.height
}

func (s square) perim() float64 {
	return 2 * (s.height + s.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geo) {
	fmt.Println("Area :", g.area(), "Perim :", g.perim())
}

func main() {
	s := square{width: 10, height: 20}
	c := circle{radius: 10}
	measure(s)
	measure(c)	
}
