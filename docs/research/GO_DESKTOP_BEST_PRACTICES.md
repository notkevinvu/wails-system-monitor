# Go Desktop Application Best Practices

> Research compiled for building a system monitor overlay with Wails

## Go Desktop vs Backend: Key Differences

### Backend Go (Traditional)
- Stateless request-response cycles
- HTTP server with handlers
- Connection pooling, middleware patterns
- Horizontal scaling focus

### Desktop Go (Local App)
| Aspect | Desktop Approach |
|--------|------------------|
| **State** | Long-lived application state in memory |
| **Lifecycle** | Startup → Running → Shutdown (user-controlled) |
| **Threading** | UI thread + background goroutines |
| **Resources** | Single user's machine resources |
| **Communication** | Frontend ↔ Backend via bindings (not HTTP) |

## Concurrency Patterns for Desktop Apps

### 1. Worker Pool Pattern
Use for CPU-intensive background tasks without blocking UI:
```go
type MetricsCollector struct {
    ctx    context.Context
    cancel context.CancelFunc
    data   chan MetricData
}

func (m *MetricsCollector) Start() {
    go m.collectLoop() // Background goroutine
}
```

### 2. Context Discipline
**Critical for desktop apps** - always propagate contexts for clean shutdown:
```go
func (a *App) Shutdown(ctx context.Context) {
    a.collector.Stop() // Signal all goroutines to stop
    // Wait for cleanup with timeout
}
```

### 3. Bounded Channels
Prefer bounded channels over unbounded queues to prevent memory leaks:
```go
metrics := make(chan CPUData, 100) // Bounded buffer
```

### 4. Fan-Out/Fan-In
Good for collecting multiple metrics simultaneously:
```go
// Fan-out: spawn collectors
go collectCPU(ctx, results)
go collectMemory(ctx, results)
go collectDisk(ctx, results)

// Fan-in: aggregate in main loop
for metric := range results { ... }
```

## System Metrics Collection (gopsutil)

### Recommended Library
[gopsutil](https://github.com/shirou/gopsutil) - Cross-platform system metrics without cgo

### Key Functions
```go
import (
    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/mem"
    "github.com/shirou/gopsutil/v3/disk"
)

// CPU usage (blocks for interval duration)
percentages, _ := cpu.Percent(time.Second, false)

// Memory stats (instant)
vmem, _ := mem.VirtualMemory()
// vmem.Total, vmem.Used, vmem.UsedPercent

// Disk usage
usage, _ := disk.Usage("/")
```

### Best Practice: Isolate Profiling
Don't mix CPU profiling with memory profiling - they interfere with each other. Collect one type at a time or use separate intervals.

## Wails-Specific Patterns

### Architecture
```
┌─────────────────────────────────────────┐
│           Wails Application             │
├──────────────────┬──────────────────────┤
│    Go Backend    │   Web Frontend       │
│  ┌────────────┐  │  ┌────────────────┐  │
│  │ App struct │◄─┼──│ wailsjs/go/... │  │
│  │  Methods   │  │  │  (auto-gen)    │  │
│  └────────────┘  │  └────────────────┘  │
│        │         │         │            │
│        ▼         │         ▼            │
│  ┌────────────┐  │  ┌────────────────┐  │
│  │  Services  │  │  │  UI Components │  │
│  │ (metrics)  │  │  │  (charts)      │  │
│  └────────────┘  │  └────────────────┘  │
└──────────────────┴──────────────────────┘
```

### Overlay Window Options (v2)
```go
wails.Run(&options.App{
    Frameless:           true,
    AlwaysOnTop:         true,
    WindowIsTranslucent: true,
    WebviewIsTransparent: true,
    // Set CSS: background-color: rgba(0,0,0,0)
})
```

### Data Streaming Pattern
```go
// Backend: emit events to frontend
func (a *App) StartMetricsStream(ctx context.Context) {
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ctx.Done():
            return
        case <-ticker.C:
            metrics := a.collector.GetLatest()
            runtime.EventsEmit(ctx, "metrics:update", metrics)
        }
    }
}
```

```javascript
// Frontend: listen for events
import { EventsOn } from '../wailsjs/runtime/runtime';

EventsOn("metrics:update", (data) => {
    updateChart(data);
});
```

## Wails v2 vs v3

| Feature | v2 (Stable) | v3 (Alpha) |
|---------|-------------|------------|
| Windows | Single | Multiple |
| System Tray | Limited | Full support |
| API Style | Declarative | Procedural |
| Stability | Production | Alpha |

**Recommendation**: Use v2 for now; v3 has better system tray but is still alpha.

## Project Structure Recommendation

```
wails-app/
├── app.go           # Main app struct + lifecycle
├── main.go          # Entry point
├── internal/
│   ├── metrics/     # System metrics collection
│   │   ├── collector.go
│   │   ├── cpu.go
│   │   ├── memory.go
│   │   └── disk.go
│   └── config/      # App configuration
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   │   └── charts/
│   │   ├── stores/   # State management
│   │   └── main.js
│   └── wailsjs/     # Auto-generated bindings
└── wails.json
```

## Sources

- [Go Best Practices 2025](https://www.bacancytechnology.com/blog/go-best-practices)
- [Go Concurrency Patterns](https://cristiancurteanu.com/7-powerful-golang-concurrency-patterns-that-will-transform-your-code-in-2025/)
- [Wails v3 What's New](https://v3alpha.wails.io/whats-new/)
- [gopsutil GitHub](https://github.com/shirou/gopsutil)
- [Building System Monitor in Go](https://rezmoss.com/blog/building-terminal-system-monitor-golang/)
- [Wails Options Reference](https://wails.io/docs/reference/options/)
- [Wails Frameless Guide](https://wails.io/docs/guides/frameless/)
