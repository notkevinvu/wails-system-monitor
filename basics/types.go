package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("--- Basic Types and Operators ---")

	// 1. Integers
	// int (platform dependent, usually 64-bit), int8, int16, int32, int64
	// uint (unsigned), uint8, etc.
	var a int = 10
	var b int = 3
	fmt.Printf("Integer Division: %d / %d = %d\n", a, b, a/b) // Result is 3
	fmt.Printf("Remainder: %d %% %d = %d\n", a, b, a%b)       // Result is 1

	// 2. Floats
	// float32, float64 (default)
	var pi float64 = 3.14159
	var radius float64 = 5.5
	area := pi * math.Pow(radius, 2)
	fmt.Printf("Float calculation (Area): %.2f\n", area)

	// 3. Booleans
	// true or false
	isGoFun := true
	isHard := false
	fmt.Println("Boolean Logic:", isGoFun && !isHard) // true AND NOT false

	// 4. Strings
	// Immutable sequence of bytes (UTF-8)
	str1 := "Hello"
	str2 := "World"
	combined := str1 + " " + str2 // Concatenation
	fmt.Println("String:", combined)
	fmt.Println("String Length:", len(combined))

	// 5. Type Conversion
	// Go requires explicit conversion
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Printf("Conversions: %d -> %f -> %d\n", i, f, u)
}
