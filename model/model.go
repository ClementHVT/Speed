package model

import (
	"main/stats"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Model represents app state
type Model struct {
	CPU float64
	Mem float64
	Err error
}

// ErrMsg wraps an error
type ErrMsg struct {
	Err error
}

// Tick returns a tea.Cmd that updates the model every second
func Tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		cpu, mem, err := stats.GetStats()
		if err != nil {
			return ErrMsg{Err: err}
		}
		return Model{CPU: cpu, Mem: mem}
	})
}

func (m Model) Init() tea.Cmd {
	return Tick()
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

	case Model:
		m.CPU = msg.CPU
		m.Mem = msg.Mem
		m.Err = nil
		return m, Tick()
	}

	return m, nil
}
