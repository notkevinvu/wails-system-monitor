# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Quick Start

**Read `docs/SESSION_SUMMARY.md` first** for current project state, recent work, and next steps.

## Project Overview

**System Monitor Overlay** - A desktop app displaying real-time CPU/memory charts as a transparent overlay.

Built with:
- **Go** + **Wails v2** (desktop framework)
- **gopsutil** (system metrics)
- **TypeScript** frontend (strict mode, async/await)

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
- **`frontend/`** - TypeScript frontend using Vite
  - `src/main.ts` - Main entry point with strict TypeScript
  - `wailsjs/go/main/App.js` - Auto-generated bindings to call Go methods
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

## Workflow Requirements

### PR Handling
- **Always include PR links**: When working with PRs, the final line(s) of any response MUST include link(s) to the PR(s) being handled
- **Feature/bugfix branches**: Always run `/full-pr-review` when working with feature or bugfix branches
- **Documentation-only changes**: Do not require full PR review workflow

### Parallel Agents
When the user requests parallel agents, immediately spawn multiple concurrent agents using the Task tool with multiple tool calls in a single message.

### Clarifying Questions
When asking clarifying questions, always use the AskUserQuestion tool to present options in the answer picker UI. Do not list questions as plain text.

### Feature Development
When using `/feature-dev`, leverage `/full-pr-review` for the PR review phase rather than basic review commands.

### Commit Strategy
Commit work in logical chunks for easier commit history digestion. Before moving to the PR review phase in `/feature-dev`, ensure all work is committed in logical, digestible chunks. Each commit should represent a coherent unit of change.

## Custom Slash Commands

Located in `.claude/commands/`:

| Command | Purpose |
|---------|---------|
| `/sync-main` | Switch to main branch and pull latest from remote |
| `/full-pr-review` | Complete PR workflow: create, review, fix, CI check, document |
| `/update-session` | Update `docs/SESSION_SUMMARY.md` with recent work and next steps |

### /full-pr-review Workflow
1. Create PR with descriptive title/summary
2. Run PR review and post findings as comment
3. Fix critical/major issues (always), minor issues (if quick or no major issues)
4. Second review cycle with fixes
5. Monitor CI and fix failures
6. Update related docs
7. Report summary with PR link(s) as final line

**Max 2 review-fix cycles** without additional prompting.

## Helper Scripts

Located in `scripts/`:

| Script | Purpose |
|--------|---------|
| `start-wails-dev.bat` | Launch Wails dev server from any location |

See `docs/HELPERS.md` for full documentation.
