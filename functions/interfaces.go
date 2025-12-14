package main

import (
	"fmt"
	"math"
)

// Define a struct
type Rect struct {
	Width, Height float64
}

// 1. Methods
// A function with a "receiver" argument.
// Value receiver: receives a copy (cannot modify original)
func (r Rect) Area() float64 {
	return r.Width * r.Height
}

// Pointer receiver: receives a pointer (can modify original)
// Generally preferred for consistency and efficiency
func (r *Rect) Scale(f float64) {
	r.Width = r.Width * f
	r.Height = r.Height * f
}

// 2. Interfaces
// A set of method signatures.
// Implicit implementation: if a type has the methods, it implements the interface.
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Function that accepts an interface
func printArea(s Shape) {
	fmt.Printf("Shape Area: %.2f\n", s.Area())
}

func main() {
	fmt.Println("--- Methods and Interfaces ---")

	r := Rect{Width: 10, Height: 5}
	fmt.Println("Rect Area:", r.Area())

	r.Scale(2)
	fmt.Printf("Scaled Rect: %+v\n", r)
	fmt.Println("New Area:", r.Area())

	c := Circle{Radius: 5}

	// Polymorphism via Interfaces
	// Both Rect and Circle implement Shape because they have an Area() method
	shapes := []Shape{r, c}
	for _, s := range shapes {
		printArea(s)
	}
}
