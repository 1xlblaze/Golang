package main

import (
	"fmt"
	"math"
)

// Define the Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

//Interface are the type of methods signature that does not provide implementation.

// Circle struct
type Circle struct {
	Radius float64
}

// Implement the Shape interface for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Implement the Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// A function that works with any type that satisfies the Shape interface
func describeShape(s Shape) {
	fmt.Printf("Shape details: Area = %.2f, Perimeter = %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	// Create a Circle and Rectangle
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 3}

	// Use polymorphism to handle different types in the same way
	describeShape(circle)    // Output: Shape details: Area = 78.54, Perimeter = 31.42
	describeShape(rectangle) // Output: Shape details: Area = 12.00, Perimeter = 14.00
}
