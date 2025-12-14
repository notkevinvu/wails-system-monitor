# Session Summary

> **Purpose**: Quick context for Claude Code to understand current project state.
> **Update**: After completing features or changing direction.

---

## Project Overview

**What**: System monitor overlay desktop app
**Stack**: Go + Wails v2 + Vanilla JS + gopsutil
**Repo**: https://github.com/notkevinvu/wails-system-monitor

---

## Recently Completed

1. **Project Setup** - Initialized git repo, created GitHub remote, established `main` branch
2. **Documentation Structure** - Created `docs/` with architecture, features, requirements, and research
3. **Go Desktop Research** - Documented best practices for Go desktop apps vs backend (see `docs/research/GO_DESKTOP_BEST_PRACTICES.md`)
4. **Project Direction** - Decided on system monitor overlay concept with real-time charts

---

## Current State

- Base Wails v2 "Hello World" scaffold exists in `wails-app/`
- Go learning examples in root folders (basics, concurrency, etc.)
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
| `wails-app/app.go` | App struct, lifecycle, bindings |
| `wails-app/frontend/src/main.js` | Frontend entry |
| `docs/ARCHITECTURE.md` | System design |
| `docs/FEATURES.md` | Feature checklist |

---

## Technical Decisions

- **Wails v2** (not v3) - v3 is alpha, v2 is stable
- **gopsutil** - Cross-platform metrics without cgo
- **Event streaming** - Push metrics via events, not polling
- **Bounded history** - Keep 60 data points max in memory

---

## Commands

```bash
# Development
cd wails-app && wails dev

# Build
cd wails-app && wails build

# Run Go examples
go run basics/variables.go
```
