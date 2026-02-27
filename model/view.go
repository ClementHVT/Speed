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

func formatBytes(bytes uint64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	switch {
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/float64(GB))
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/float64(MB))
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/float64(KB))
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}

// View renders the terminal UI
func (m Model) View() string {
	if m.Err != nil {
		return fmt.Sprintf("Error: %v\n\nRetrying...\n\nPress Ctrl+C to quit", m.Err)
	}

	filledCPU := int(m.CPU / 5)
	filledMem := int(m.Mem / 5)

	return fmt.Sprintf(
		"%s CPU : %.2f%%\n\n%s Memory : %.2f%%  %s/%s\n\nPress Ctrl+C to quit.",
		m.UsageBar(filledCPU),
		m.CPU,
		m.UsageBar(filledMem),
		m.Mem,
		formatBytes(m.MemUsed),
		formatBytes(m.MemTotal),
	)
}
