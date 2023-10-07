package main

import (
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
	"tic-tac-toe/board"
)

type model struct {
	board board.Board

	width  int
	height int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	return m.board.View()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		msg.Width -= 4
		msg.Height -= 2
		return m, nil
	}
	b, _ := m.board.Update(msg)
	m.board = b.(board.Board)
	return m, nil
}

func New() model {
	return model{
		board:  board.NewBoard(),
		width:  0,
		height: 0,
	}
}

func main() {
	zone.NewGlobal()
	p := tea.NewProgram(New())
	if err, _ := p.Run(); err != nil {
		panic(err)
	}
}
