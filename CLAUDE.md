# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Quick Start

**Read `docs/SESSION_SUMMARY.md` first** for current project state, recent work, and next steps.

## Project Overview

**System Monitor Overlay** - A desktop app displaying real-time CPU/memory charts as a transparent overlay.

Built with:
- **Go** + **Wails v2** (desktop framework)
- **gopsutil** (system metrics)
- **Vanilla JS** frontend with charts

### Repository Structure
1. **Go language examples** - Learning files in root (`basics/`, `concurrency/`, etc.)
2. **Wails desktop application** - Main app in `wails-app/`
3. **Documentation** - Architecture, features, research in `docs/`

## Commands

### Running Go Examples
```bash
go run basics/variables.go
go run controlflow/loops.go
go run concurrency/channels.go
```

### Wails Application (in `wails-app/` directory)
```bash
# Development mode with hot reload
wails dev

# Production build
wails build

# Frontend dev server accessible at http://localhost:34115
```

### Frontend Only (in `wails-app/frontend/` directory)
```bash
npm install
npm run dev      # Vite dev server
npm run build    # Production build
```

## Architecture

### Go Examples Structure
Each folder contains standalone executable Go files demonstrating specific concepts:
- `basics/` - Variables, types, constants
- `controlflow/` - Conditionals, loops
- `datastructures/` - Slices, maps, structs
- `functions/` - Functions, methods, interfaces
- `concurrency/` - Goroutines, channels

### Wails Application Architecture
The Wails app follows the standard Wails v2 pattern:

- **`main.go`** - Application entry point, configures Wails options, embeds frontend assets via `//go:embed`
- **`app.go`** - Contains the `App` struct with methods exposed to the frontend via `Bind`
- **`frontend/`** - Vanilla JS frontend using Vite
  - `wailsjs/go/main/App.js` - Auto-generated bindings to call Go methods from JS
  - `wailsjs/runtime/` - Wails runtime for frontend-backend communication

### Go-Frontend Binding
Methods on the `App` struct bound in `main.go` become callable from JavaScript:
```go
// Go (app.go)
func (a *App) Greet(name string) string { ... }
```
```javascript
// JavaScript
import { Greet } from '../wailsjs/go/main/App';
const result = await Greet("World");
```
