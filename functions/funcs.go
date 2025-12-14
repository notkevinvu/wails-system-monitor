package main

import "fmt"

// 1. Basic Function
func add(a int, b int) int {
	return a + b
}

// 2. Multiple Return Values
// A key feature of Go (often used for returning result + error)
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// 3. Named Return Values
// Variables 'min' and 'max' are initialized automatically
func minMax(a, b int) (min, max int) {
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}
	return // Naked return (returns min, max)
}

// 4. Variadic Functions
// Accepts any number of arguments
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println("--- Functions ---")

	// Basic
	fmt.Println("Add:", add(5, 3))

	// Multiple Returns
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division:", result)
	}

	// Error handling
	_, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println("Error handling works:", err2)
	}

	// Named Returns
	min, max := minMax(10, 5)
	fmt.Printf("Min: %d, Max: %d\n", min, max)

	// Variadic
	fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))
	fmt.Println("Sum of nothing:", sum())
}
