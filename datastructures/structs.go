package main

import "fmt"

// 1. Defining a Struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
}

// Structs can be nested
type Employee struct {
	Details Person
	JobID   string
	Salary  int
}

func main() {
	fmt.Println("--- Structs ---")

	// 2. Creating Instances
	// Named fields (Recommended)
	p1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       40,
		Email:     "john@example.com",
	}
	fmt.Printf("Person 1: %+v\n", p1) // %+v prints field names too

	// Omitted fields get zero values
	p2 := Person{
		FirstName: "Jane",
	}
	fmt.Printf("Person 2: %+v\n", p2)

	// 3. Accessing Fields
	fmt.Println("P1 Email:", p1.Email)
	p1.Age = 41
	fmt.Println("P1 New Age:", p1.Age)

	// 4. Anonymous Structs
	// Useful for one-off data structures (e.g. JSON parsing)
	config := struct {
		Port int
		Env  string
	}{
		Port: 8080,
		Env:  "Production",
	}
	fmt.Println("Config:", config)

	// 5. Pointers to Structs
	// Efficiently pass large structs around without copying
	p3 := &Person{FirstName: "Pointer", Age: 20}
	fmt.Println("Pointer to Person:", p3)
	// Go automatically dereferences pointers for struct fields
	fmt.Println("Accessing via pointer:", p3.FirstName)
}
