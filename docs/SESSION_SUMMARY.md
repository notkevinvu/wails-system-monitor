# Session Summary

> **Purpose**: Quick context for Claude Code to understand current project state.
> **Update**: After completing features or changing direction.

---

## Project Overview

**What**: System monitor overlay desktop app
**Stack**: Go + Wails v2 + TypeScript + gopsutil
**Repo**: https://github.com/notkevinvu/wails-system-monitor

---

## Recently Completed

1. **Go-JS Data Binding Demo** - Added `GetItems()` function in Go with frontend button to fetch/display list
2. **Claude Workflow Commands** - Created `/sync-main`, `/full-pr-review`, `/update-session` slash commands
3. **Developer Experience** - Added `scripts/start-wails-dev.bat` for easy dev server launch from anywhere

---

## Current State

- Wails v2 app with working Go-JS data binding (`Greet`, `GetItems`)
- Go learning examples in root folders (basics, concurrency, etc.)
- Claude workflow automation via custom slash commands
- No metrics collection implemented yet
- No overlay window configuration yet

---

## Next Steps

1. **Add gopsutil dependency** - `go get github.com/shirou/gopsutil/v3` in wails-app
2. **Create metrics collector** - Background goroutine collecting CPU/memory at 1s intervals
3. **Wire up event streaming** - Use `runtime.EventsEmit` to push metrics to frontend
4. **Configure overlay window** - Set frameless, always-on-top, transparent options in `main.go`

---

## Key Files

| File | Purpose |
|------|---------|
| `wails-app/main.go` | App entry, Wails options |
| `wails-app/app.go` | App struct, lifecycle, bindings (Greet, GetItems) |
| `wails-app/frontend/src/main.ts` | Frontend entry, UI building (TypeScript) |
| `scripts/start-wails-dev.bat` | Launch dev server from anywhere |
| `.claude/commands/*.md` | Custom slash commands |
| `docs/HELPERS.md` | Helper scripts documentation |
| `docs/ARCHITECTURE.md` | System design |
| `docs/FEATURES.md` | Feature checklist |

---

## Technical Decisions

- **Wails v2** (not v3) - v3 is alpha, v2 is stable
- **TypeScript** (strict mode) - Type-safe frontend with async/await
- **gopsutil** - Cross-platform metrics without cgo
- **Event streaming** - Push metrics via events, not polling
- **Bounded history** - Keep 60 data points max in memory

---

## Commands

```bash
# Development (easy way)
scripts\start-wails-dev.bat

# Development (manual)
cd wails-app && wails dev

# Build
cd wails-app && wails build

# Run Go examples
go run basics/variables.go
```

### Claude Slash Commands

| Command | Purpose |
|---------|---------|
| `/sync-main` | Switch to main and pull latest |
| `/full-pr-review` | Complete PR workflow with review cycles |
| `/update-session` | Update this SESSION_SUMMARY.md |
