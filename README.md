# Go and Wails Learning Project

This project contains a structured curriculum for learning Go (Golang) and an introduction to the Wails framework.

## Curriculum Modules

### 1. Language Basics
- **[basics/variables.go](basics/variables.go)**: Variables, constants, and short declarations.
- **[basics/types.go](basics/types.go)**: Basic types (int, float, bool, string) and operators.

### 2. Control Flow
- **[controlflow/conditionals.go](controlflow/conditionals.go)**: `if/else` and `switch` statements.
- **[controlflow/loops.go](controlflow/loops.go)**: `for` loops (the only loop in Go) and `range`.

### 3. Data Structures
- **[datastructures/slices.go](datastructures/slices.go)**: Arrays and Slices (dynamic arrays).
- **[datastructures/maps.go](datastructures/maps.go)**: Key-value stores.
- **[datastructures/structs.go](datastructures/structs.go)**: Custom data types and fields.

### 4. Functions and Methods
- **[functions/funcs.go](functions/funcs.go)**: Function syntax, multiple returns, named returns.
- **[functions/interfaces.go](functions/interfaces.go)**: Methods on structs and Interfaces.

### 5. Concurrency
- **[concurrency/goroutines.go](concurrency/goroutines.go)**: Lightweight threads (`go` keyword) and `sync.WaitGroup`.
- **[concurrency/channels.go](concurrency/channels.go)**: Communication between goroutines.

### 6. Wails Application
- **[wails-app/](wails-app/)**: A "Hello World" desktop application built with Wails (Go backend + Web frontend).

## How to Run

### Running Go Examples
You can run any individual file using `go run`:

```bash
go run basics/variables.go
go run concurrency/channels.go
```

### Running the Wails App
To start the Wails desktop application in development mode (with hot reload):

```bash
cd wails-app
wails dev
```

This will compile the Go backend and serve the frontend, opening a window with your application.
