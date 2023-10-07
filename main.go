package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/lrstanley/bubblezone"
	"log"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func main() {
	zone.NewGlobal()
	p := tea.NewProgram(
		New(),
		tea.WithAltScreen(),
		//tea.WithMouseCellMotion(),
		tea.WithMouseAllMotion(),
	)

	if err, _ := p.Run(); err != nil {
		log.Fatal(err)
	}
}
