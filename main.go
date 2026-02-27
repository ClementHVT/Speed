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
	err error
}

type errMsg struct {
	err error
}

func getStats() (float64, float64, error) {
	percent, err := cpu.Percent(0, false)

	if err != nil {
		return 0, 0, err
	}

	if len(percent) == 0 {
		return 0, 0, fmt.Errorf("no CPU data returned")
	}
	v, _ := mem.VirtualMemory()

	if err != nil {
		return 0, 0, err
	}

	return percent[0], v.UsedPercent, nil
}

func (m model) Init() tea.Cmd {
	return tick()
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		cpu, mem, err := getStats()
		if err != nil {
			return errMsg{err}
		}
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

	case errMsg:
		m.err = msg.err
		return m, tick()

	case model:
		m.cpu = msg.cpu
		m.mem = msg.mem
		m.err = nil
		return m, tick()
	}
	return m, nil
}

func (m model) UsageBar(filledValue int) string {
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

func (m model) View() string {

	if m.err != nil {
		return fmt.Sprintf(
			"Error: %v\n\n Retrying...\n\nPress Ctrl+C to quit",
			m.err,
		)
	}

	filledCpu := int(m.cpu / 5)
	filledMem := int(m.mem / 5)

	cpu := m.UsageBar(filledCpu)
	mem := m.UsageBar(filledMem)

	return fmt.Sprintf(
		"%s CPU : %.2f%%\n\n%s Memory : %.2f%%\n\nPress Ctrl+C to quit.",
		cpu,
		m.cpu,
		mem,
		m.mem,
	)
}

func main() {
	p := tea.NewProgram(model{})
	if err := p.Start(); err != nil {
		fmt.Println("Error:", err)
	}
}
