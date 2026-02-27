package main

import (
	"fmt"
	"main/model"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(model.Model{})
	if err := p.Start(); err != nil {
		fmt.Println("Error:", err)
	}
}
