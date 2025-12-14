package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("--- Conditionals ---")

	// 1. If / Else
	// Parentheses are not required around the condition, but braces are required.
	age := 18
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// 2. If with Short Statement
	// You can execute a statement before the condition.
	// The variable 'v' is only available inside the if/else block.
	if v := 10; v > 5 {
		fmt.Println("v is greater than 5")
	}

	// 3. Switch
	// No 'break' needed! Go breaks automatically.
	// Use 'fallthrough' to force execution of the next case.
	os := runtime.GOOS
	fmt.Print("Go runs on ")
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	case "windows":
		fmt.Println("Windows.")
	default:
		fmt.Printf("%s.\n", os)
	}

	// 4. Switch with no condition (cleaner if/else chain)
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
