# System Monitor Overlay - Requirements

## Functional Requirements

### FR-1: System Metrics Collection
- **FR-1.1**: Application SHALL collect CPU usage percentage at configurable intervals
- **FR-1.2**: Application SHALL collect memory usage (total, used, percentage) at configurable intervals
- **FR-1.3**: Default collection interval SHALL be 1 second
- **FR-1.4**: Collection SHALL run in a background goroutine without blocking UI

### FR-2: Data Visualization
- **FR-2.1**: Application SHALL display metrics as real-time updating charts
- **FR-2.2**: Charts SHALL show rolling history (minimum 60 data points)
- **FR-2.3**: Charts SHALL update smoothly without flickering

### FR-3: Overlay Window
- **FR-3.1**: Window SHALL be frameless (no OS title bar)
- **FR-3.2**: Window SHALL support always-on-top mode
- **FR-3.3**: Window SHALL be repositionable via drag gesture
- **FR-3.4**: Window background SHALL support transparency/translucency

### FR-4: Application Lifecycle
- **FR-4.1**: Application SHALL start metrics collection on launch
- **FR-4.2**: Application SHALL gracefully stop all goroutines on close
- **FR-4.3**: Application SHALL not leak resources on shutdown

---

## Non-Functional Requirements

### NFR-1: Performance
- **NFR-1.1**: Idle CPU usage SHALL be < 2%
- **NFR-1.2**: Memory footprint SHALL be < 100MB
- **NFR-1.3**: UI SHALL remain responsive during metric collection

### NFR-2: Compatibility
- **NFR-2.1**: Application SHALL run on Windows 10/11
- **NFR-2.2**: Application SHOULD run on macOS 11+ (best effort)
- **NFR-2.3**: Application SHOULD run on Linux with GTK (best effort)

### NFR-3: Usability
- **NFR-3.1**: Application SHALL be usable without documentation
- **NFR-3.2**: Charts SHALL be readable at default window size
- **NFR-3.3**: Window SHALL remember position between sessions (Phase 2)

### NFR-4: Code Quality
- **NFR-4.1**: Go code SHALL pass `go vet` without warnings
- **NFR-4.2**: Go code SHALL be formatted with `gofmt`
- **NFR-4.3**: Critical paths SHALL have unit tests

---

## Constraints

### Technical Constraints
- Must use Wails v2 framework (learning objective)
- Must use gopsutil for cross-platform metrics
- Frontend limited to vanilla JS (no React/Vue for simplicity)

### Project Constraints
- Learning project - prioritize understanding over perfection
- Single developer
- No external dependencies beyond Go stdlib + Wails + gopsutil + charting lib

---

## Assumptions

1. User has Go 1.21+ installed
2. User has Wails CLI installed (`go install github.com/wailsapp/wails/v2/cmd/wails@latest`)
3. User has Node.js 18+ for frontend tooling
4. Target machines have standard system APIs accessible (no sandboxed environments)
