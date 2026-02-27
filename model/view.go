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

	filledCPU := int(m.CPU / 5)
	filledMem := int(m.Mem / 5)

	return fmt.Sprintf(
		"%s CPU : %.2f%%\n\n%s Memory : %.2f%%\n\nPress Ctrl+C to quit.",
		m.UsageBar(filledCPU),
		m.CPU,
		m.UsageBar(filledMem),
		m.Mem,
	)
}
