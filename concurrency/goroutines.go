package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. Goroutines
// A goroutine is a lightweight thread managed by the Go runtime.
// You start one by adding the 'go' keyword before a function call.

func printMessage(msg string, count int, wg *sync.WaitGroup) {
	// Defer ensuring Done() is called when function exits
	defer wg.Done()

	for i := 0; i < count; i++ {
		fmt.Printf("%s: %d\n", msg, i)
		time.Sleep(100 * time.Millisecond) // Simulate work
	}
}

func main() {
	fmt.Println("--- Goroutines ---")

	// 2. Synchronization
	// Since goroutines run concurrently, main() might finish before they do.
	// We use sync.WaitGroup to wait for them.
	var wg sync.WaitGroup

	// We are launching 2 goroutines
	wg.Add(2)

	// Launch goroutines
	go printMessage("Goroutine A", 3, &wg)
	go printMessage("Goroutine B", 3, &wg)

	fmt.Println("Main: Waiting for goroutines to finish...")
	wg.Wait() // Blocks until counter is 0
	fmt.Println("Main: All done!")
}
