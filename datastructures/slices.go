package main

import "fmt"

func main() {
	fmt.Println("--- Arrays and Slices ---")

	// 1. Arrays
	// Fixed size. Not commonly used directly in Go logic, but slices are built on them.
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println("Array:", arr)

	// 2. Slices
	// Dynamic size. This is what you use 99% of the time.
	// Syntax: []T (no size specified)
	slice := []string{"a", "b", "c"}
	fmt.Println("Slice:", slice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice), cap(slice))

	// 3. Appending
	// append returns a NEW slice (because the underlying array might change)
	slice = append(slice, "d")
	slice = append(slice, "e", "f")
	fmt.Println("Appended slice:", slice)

	// 4. Slicing (Sub-slices)
	// slice[low : high] (inclusive low, exclusive high)
	// indices: 0="a", 1="b", 2="c", 3="d", 4="e", 5="f"
	subSlice := slice[1:4] // "b", "c", "d"
	fmt.Println("Sub-slice [1:4]:", subSlice)

	// 5. Making Slices
	// make([]T, len, cap)
	// Useful to pre-allocate memory if you know the size.
	numbers := make([]int, 5) // [0 0 0 0 0]
	numbers[0] = 100
	fmt.Println("Made slice:", numbers)
}
