package main

import "fmt"

// Shape interface
type shape interface {
	getArea() float64
}

// Triangle type
type triangle struct {
	base   float64
	height float64
}

// Square type
type square struct {
	sideLength float64
}

func main() {
	t := triangle{
		base:   3,
		height: 5,
	}

	s := square{
		sideLength: 6,
	}

	fmt.Println("Area of the triangle:")
	printArea(t)
	fmt.Println("Area of the square:")
	printArea(s)
}

// Method related to the shape interface which
// prints the shape area
func printArea(s shape) {
	fmt.Println(s.getArea())
}

// Methods that returns the area of a triangle
func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

// Methods that returns the area of a square
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
