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

	return fmt.Sprintf(
		"%s CPU : %.2f%%\n\n%s Memory : %.2f%%  %s/%s\n\nPress Ctrl+C to quit.",
		m.UsageBar(m.UI.CPUBarFill),
		m.UI.CPUPercent,
		m.UsageBar(m.UI.MemBarFill),
		m.UI.MemPercent,
		m.UI.MemUsedStr,
		m.UI.MemTotalStr,
	)
}
