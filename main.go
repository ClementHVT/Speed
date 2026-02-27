package main

import (
	"fmt"
	"main/model"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var version = "dev"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Printf("Speed Version: %s\n", version)
		return
	}

	p := tea.NewProgram(model.Model{
		Version: version,
	})

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
