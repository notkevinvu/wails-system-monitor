package main

import "fmt"

func main() {
	fmt.Println("--- Variables and Constants ---")

	// 1. Standard Declaration (var name type)
	var age int = 30
	fmt.Println("Standard var:", age)

	// 2. Type Inference (var name = value)
	var name = "Alice" // Go infers this is a string
	fmt.Println("Inferred var:", name)

	// 3. Short Declaration (name := value)
	// Only works inside functions. Most common way to declare variables.
	country := "Canada"
	fmt.Println("Short declaration:", country)

	// 4. Multiple Declaration
	var x, y int = 10, 20
	fmt.Println("Multiple vars:", x, y)

	// 5. Zero Values
	// Variables declared without a value get a "zero value"
	var z int       // 0
	var s string    // ""
	var b bool      // false
	fmt.Printf("Zero values: int=%d, string=%q, bool=%v\n", z, s, b)

	// 6. Constants
	// Cannot be changed. Can be character, string, boolean, or numeric values.
	const Pi = 3.14159
	const Greeting = "Hello"
	fmt.Println("Constant:", Pi)
	// Pi = 3.14 // This would cause a compile error
}
