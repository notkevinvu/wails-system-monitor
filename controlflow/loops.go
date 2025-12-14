package main

import "fmt"

func main() {
	fmt.Println("--- Loops ---")

	// 1. Basic For Loop (C-style)
	// The only loop construct in Go is 'for'.
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Println("Sum (0..4):", sum)

	// 2. While-style Loop
	// Omit the init and post statements.
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println("While-style result:", n)

	// 3. Infinite Loop
	// for { ... }
	count := 0
	for {
		count++
		if count > 3 {
			break // Exit the loop
		}
	}
	fmt.Println("Exited infinite loop")

	// 4. Range Loop (Iterating over collections)
	// Used for slices, arrays, maps, strings, and channels.
	nums := []int{2, 4, 6}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// If you only need the value, use _ to ignore the index
	for _, value := range nums {
		fmt.Printf("Value only: %d\n", value)
	}
}
