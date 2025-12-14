package main

import "fmt"

func main() {
	fmt.Println("--- Maps ---")

	// 1. Creating a Map
	// map[KeyType]ValueType
	colors := make(map[string]string)
	colors["red"] = "#ff0000"
	colors["green"] = "#00ff00"
	colors["blue"] = "#0000ff"

	fmt.Println("Colors map:", colors)

	// 2. Map Literal
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	fmt.Println("Ages map:", ages)

	// 3. Accessing and Checking Existence
	// If a key doesn't exist, you get the zero value.
	// To know if it really exists, use the second return value (ok).
	age, exists := ages["Charlie"]
	if exists {
		fmt.Println("Charlie is", age)
	} else {
		fmt.Println("Charlie is not in the map")
	}

	// 4. Deleting
	delete(ages, "Bob")
	fmt.Println("After deleting Bob:", ages)

	// 5. Iterating
	// Order is random!
	for key, value := range colors {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}
