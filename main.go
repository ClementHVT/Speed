package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cpu float64
	mem float64
}

func getStats() (float64, float64) {
	percent, _ := cpu.Percent(0, false)
	v, _ := mem.VirtualMemory()

	return percent[0], v.UsedPercent
}

func (m model) Init() tea.Cmd {
	return tick()
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		cpu, mem := getStats()
		return model{cpu: cpu, mem: mem}
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case model:
		m.cpu = msg.cpu
		m.mem = msg.mem
		return m, tick()
	}
	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf(
		"CPU Usage: %.2f%%\nMemory Usage: %.2f%%\n\nPress Ctrl+C to quit.",
		m.cpu,
		m.mem,
	)
}

func main() {
	p := tea.NewProgram(model{})
	if err := p.Start(); err != nil {
		fmt.Println("Error:", err)
	}
}
