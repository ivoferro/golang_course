package main

import "fmt"

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println("Area:", s.getArea())
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return t.height * t.base / 2
}

func main() {
	s := square{sideLength: 5}
	printArea(s)

	t := triangle{base: 5, height: 4}
	printArea(t)
}
