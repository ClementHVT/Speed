package model

import "fmt"

// UsageBar returns a simple bar graph for percentage values
func (m Model) UsageBar(filledValue int) string {
	total := 20
	bar := "[ "

	for i := 0; i < total; i++ {
		if i < filledValue {
			bar += "| "
		} else {
			bar += "  "
		}
	}

	bar += "]"
	return bar
}

// View renders the terminal UI
func (m Model) View() string {
	if m.Err != nil {
		return fmt.Sprintf("Error: %v\n\nRetrying...\n\nPress Ctrl+C to quit", m.Err)
	}

	// Header
	header := fmt.Sprintf("─── SPEED by ClementHVT [%s] ──────────────────────────", m.Version)
	view := header + "\n\n"

	// CPU Average
	view += fmt.Sprintf(
		"CPU (Avg: %.1f%%) %s %.1f%%\n",
		m.UI.CPUPercent,
		m.UsageBar(m.UI.CPUBarFill),
		m.UI.CPUPercent,
	)

	// Per-Core Usage
	view += "Per-Core Usage:\n"
	for i, usage := range m.CPU.PerCoreUsage {
		barFill := int(usage / 5)
		view += fmt.Sprintf("  Core %02d: %s %.1f%%\n", i, m.UsageBar(barFill), usage)
	}

	view += "\n"

	// Memory
	view += fmt.Sprintf(
		"Memory: %.1f%% %s %s/%s\n",
		m.UI.MemPercent,
		m.UsageBar(m.UI.MemBarFill),
		m.UI.MemUsedStr,
		m.UI.MemTotalStr,
	)

	// Disk
	view += fmt.Sprintf(
		"Disk: %.1f%% %s %s/%s\n",
		m.UI.DiskPercent,
		m.UsageBar(int(m.UI.DiskPercent/5)),
		m.UI.DiskUsedStr,
		m.UI.DiskTotalStr,
	)

	view += "\n───────────────────────────────────────────────────────────────\n"
	view += "Press Ctrl+C to quit"

	return view
}
