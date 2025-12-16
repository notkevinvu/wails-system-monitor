# Helper Scripts

Quick-launch scripts for common development tasks.

---

## Available Scripts

All scripts are located in the `scripts/` folder.

### `start-wails-dev.bat`

**Purpose**: Start the Wails development server with hot-reload.

**Usage**: Double-click or run from any terminal location:
```cmd
D:\Projects-Coding\GoWailsLearning\scripts\start-wails-dev.bat
```

**What it does**:
1. Changes to the `wails-app/` directory
2. Runs `wails dev` to start the development server
3. Opens the app window automatically
4. Frontend changes hot-reload, Go changes trigger rebuild

**To stop**: Press `Ctrl+C` in the terminal.

---

## Adding New Scripts

When adding new helper scripts:
1. Place them in the `scripts/` folder
2. Use absolute paths so they work from any location
3. Document them in this file
