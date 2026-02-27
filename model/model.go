package model

import (
	"fmt"
	"main/stats"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Model represents app state
type Model struct {
	CPU stats.CPUStats
	Mem stats.MemStats
	Err error

	UI UIData
}

type UIData struct {
	CPUPercent  float64
	MemPercent  float64
	CPUBarFill  int
	MemBarFill  int
	MemUsedStr  string
	MemTotalStr string
}

// ErrMsg wraps an error
type ErrMsg struct {
	Err error
}

type StatsMsg struct {
	CPU stats.CPUStats
	Mem stats.MemStats
}

// Tick returns a tea.Cmd that updates the model every second
func Tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		cpu, mem, err := stats.GetStats()
		if err != nil {
			return ErrMsg{Err: err}
		}
		return StatsMsg{CPU: cpu, Mem: mem}
	})
}

func (m Model) Init() tea.Cmd {
	return Tick()
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

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case ErrMsg:
		m.Err = msg.Err
		return m, Tick()

	case StatsMsg:
		m.CPU = msg.CPU
		m.Mem = msg.Mem
		m.Err = nil

		m.UI.CPUPercent = m.CPU.Usage
		m.UI.MemPercent = m.Mem.UsedPercent
		m.UI.CPUBarFill = int(m.UI.CPUPercent / 5)
		m.UI.MemBarFill = int(m.UI.MemPercent / 5)
		m.UI.MemUsedStr = formatBytes(m.Mem.Used)
		m.UI.MemTotalStr = formatBytes(m.Mem.Total)

		return m, Tick()
	}

	return m, nil
}
