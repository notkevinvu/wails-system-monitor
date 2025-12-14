# System Monitor Overlay - Architecture

## Overview

A lightweight desktop overlay application that displays real-time system metrics (CPU, memory, disk, network) using Go + Wails.

## High-Level Architecture

```
┌────────────────────────────────────────────────────────────────┐
│                    System Monitor Overlay                       │
├────────────────────────────────┬───────────────────────────────┤
│         Go Backend             │        Web Frontend            │
│                                │                                │
│  ┌──────────────────────────┐  │  ┌──────────────────────────┐  │
│  │        App Struct        │  │  │      UI Layer            │  │
│  │  - Lifecycle management  │◄─┼──│  - Charts (Chart.js?)    │  │
│  │  - Frontend bindings     │  │  │  - Settings panel        │  │
│  │  - Event emission        │──┼─►│  - Drag handle           │  │
│  └──────────────────────────┘  │  └──────────────────────────┘  │
│             │                  │             │                  │
│             ▼                  │             ▼                  │
│  ┌──────────────────────────┐  │  ┌──────────────────────────┐  │
│  │    Metrics Collector     │  │  │    State Management      │  │
│  │  - CPU (gopsutil)        │  │  │  - Metrics history       │  │
│  │  - Memory                │  │  │  - User preferences      │  │
│  │  - Disk                  │  │  │  - Chart config          │  │
│  │  - Network (optional)    │  │  └──────────────────────────┘  │
│  └──────────────────────────┘  │                                │
│             │                  │                                │
│             ▼                  │                                │
│  ┌──────────────────────────┐  │                                │
│  │    Config Manager        │  │                                │
│  │  - Window position       │  │                                │
│  │  - Update interval       │  │                                │
│  │  - Displayed metrics     │  │                                │
│  └──────────────────────────┘  │                                │
└────────────────────────────────┴───────────────────────────────┘
```

## Data Flow

```
┌─────────┐     ┌───────────────┐     ┌──────────────┐     ┌─────────┐
│ System  │────►│ gopsutil      │────►│ Collector    │────►│ Events  │
│ APIs    │     │ (abstraction) │     │ (goroutine)  │     │ Emit    │
└─────────┘     └───────────────┘     └──────────────┘     └────┬────┘
                                                                │
                    ┌───────────────────────────────────────────┘
                    │
                    ▼
             ┌──────────────┐     ┌──────────────┐     ┌──────────┐
             │ EventsOn     │────►│ Update State │────►│ Re-render│
             │ (frontend)   │     │ (JS)         │     │ Charts   │
             └──────────────┘     └──────────────┘     └──────────┘
```

## Component Responsibilities

### Go Backend

| Component | Responsibility |
|-----------|---------------|
| `App` | Application lifecycle, window management, frontend bindings |
| `MetricsCollector` | Background goroutine collecting system metrics at intervals |
| `ConfigManager` | Persist/load user preferences (JSON file) |

### Frontend

| Component | Responsibility |
|-----------|---------------|
| Charts | Render real-time line/area charts for metrics |
| Settings | UI for configuring update interval, visible metrics, opacity |
| Drag Handle | Allow repositioning frameless window |

## Key Design Decisions

### 1. Event-Based Communication
Use Wails runtime events (`EventsEmit`/`EventsOn`) for streaming metrics rather than polling from frontend. This is more efficient and provides real-time updates.

### 2. Single Collector Goroutine
One background goroutine handles all metric collection on a timer. Avoids spawning multiple goroutines and simplifies lifecycle management.

### 3. Bounded History
Keep only last N data points in memory (e.g., 60 for 1-minute history at 1s intervals). Prevents unbounded memory growth.

### 4. Frameless Overlay
Use Wails frameless + always-on-top + transparent options for overlay appearance. Custom CSS drag handle for repositioning.

## Technology Choices

| Layer | Technology | Rationale |
|-------|------------|-----------|
| Backend | Go 1.21+ | Learning project, good concurrency |
| Framework | Wails v2 | Stable, good desktop integration |
| Metrics | gopsutil | Cross-platform, no cgo |
| Frontend | Vanilla JS | Keep it simple, match existing scaffold |
| Charts | TBD (Chart.js or uPlot) | Lightweight, real-time capable |

## File Structure (Planned)

```
wails-app/
├── app.go                    # App struct + lifecycle hooks
├── main.go                   # Entry point + Wails config
├── internal/
│   ├── metrics/
│   │   ├── collector.go      # Main collector with goroutine
│   │   ├── types.go          # Metric data structures
│   │   └── system.go         # gopsutil wrappers
│   └── config/
│       └── config.go         # User preferences
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   │   ├── Chart.js      # Chart component
│   │   │   └── Settings.js   # Settings panel
│   │   ├── main.js           # App entry, event listeners
│   │   └── style.css         # Overlay styling
│   └── wailsjs/              # Auto-generated
└── wails.json
```
