package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

// In Go, you cannot write the printArea function with a receiver of type shape
// because Go does not allow you to define methods on interfaces directly.
// Methods in Go are associated with named types, not interfaces.
// When you define a method with a receiver, t
// he receiver must be a named type (struct or non-interface type).
// The reason behind this design is that Go interfaces do not have any fields,
// thus, it is not possible to implement methods on them directly.

func main() {
	t := triangle{2, 3}
	s := square{3}
	printArea(t)
	printArea(s)
}
