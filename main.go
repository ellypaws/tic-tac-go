package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lrstanley/bubblezone"
	"log"
	"math/rand"
	"strings"
	"tic-tac-toe/board"
	"tic-tac-toe/help"
)

type Model struct {
	Board         *board.Board
	Current       board.Cell
	GameOver      bool
	ShowError     bool
	Cursor        struct{ Row, Col int }
	Width, Height int
	help          help.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		msg.Height -= 2
		msg.Width -= 4
		return m, nil
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}

		switch {
		case key.Matches(msg, m.help.Keys.Reset):
			newModel := New()
			newModel.Width = m.Width
			newModel.Height = m.Height
			return &newModel, nil
		}
		switch msg.Type {
		case tea.KeyUp:
			m.Cursor.Col = max(m.Cursor.Col-1, 0)
		case tea.KeyDown:
			m.Cursor.Col = min(m.Cursor.Col+1, 2)
		case tea.KeyLeft:
			m.Cursor.Row = max(m.Cursor.Row-1, 0)
		case tea.KeyRight:
			m.Cursor.Row = min(m.Cursor.Row+1, 2)
		case tea.KeyEnter:
			if m.Board.Cells[m.Cursor.Col][m.Cursor.Row] == board.Empty {
				m.Board.Cells[m.Cursor.Col][m.Cursor.Row] = m.Current
				if m.Current == board.PlayerX {
					m.Current = board.PlayerO
				} else {
					m.Current = board.PlayerX
				}
			} else {
				m.ShowError = true
				return m, nil
			}
		}
		m.help, _ = m.help.Update(msg)
		return m, nil
	case tea.MouseMsg:
		for row := 0; row < 3; row++ {
			for col := 0; col < 3; col++ {
				z := zone.Get(m.Board.Id + fmt.Sprintf("%d-%d", row, col))
				if z.InBounds(msg) {
					m.Cursor.Row = row
					m.Cursor.Col = col
				}
			}
		}
		winner := m.Board.Winner()
		if winner != board.Empty || m.Board.IsFull() {
			m.GameOver = true
			return m, nil
		}
		if msg.Type == tea.MouseLeft {
			if m.Board.Cells[m.Cursor.Col][m.Cursor.Row] == board.Empty {
				m.Board.Cells[m.Cursor.Col][m.Cursor.Row] = m.Current
				if m.Current == board.PlayerX {
					m.Current = board.PlayerO
				} else {
					m.Current = board.PlayerX
				}
			} else {
				m.ShowError = true
				return m, nil
			}
		}
		return m, nil
	}
	return m, nil
}

//const (
//	EmptySymbol = "   "
//	XSymbol     = " ❌ "
//	OSymbol     = " ⭕ "
//)

// Nerd Fonts
//const (
//	EmptySymbol = " "
//	XSymbol     = "\uF467"
//	OSymbol     = "\uEABC"
//)

const (
	EmptySymbol = "   "
	XSymbol     = " X "
	OSymbol     = " O "
)

func (m Model) View() string {
	var out strings.Builder

	normalStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("228")).
		Bold(true).
		Foreground(lipgloss.Color("#ff0000")).
		Padding(1, 1)

	cursorStyle := normalStyle.Copy().
		BorderForeground(lipgloss.Color("86"))

	out.WriteString(fmt.Sprintf("[%v]:[%v] {%v:%v}\n", m.Cursor.Row, m.Cursor.Col, m.Width, m.Height))

	for col := 0; col < 3; col++ {
		var rowItems []string
		for row := 0; row < 3; row++ {
			//zoneID := fmt.Sprintf("%d-%d", row, col)

			var r = normalStyle.Render
			if col == m.Cursor.Col && row == m.Cursor.Row {
				r = cursorStyle.Render
			}

			var cell = r(EmptySymbol)
			//
			switch m.Board.Cells[col][row] {
			case board.PlayerX:
				cell = r(XSymbol)
			case board.PlayerO:
				cell = r(OSymbol)
			}

			cell = zone.Mark(m.Board.Id+fmt.Sprintf("%d-%d", row, col), cell)
			rowItems = append(rowItems, cell)
		}
		out.WriteString(lipgloss.JoinHorizontal(0.1, rowItems...))
		out.WriteString("\n")
	}

	if m.GameOver {
		winner := m.Board.Winner()
		switch winner {
		case board.PlayerX:
			out.WriteString("Player X wins!\n")
		case board.PlayerO:
			out.WriteString("Player O wins!\n")
		default:
			out.WriteString("It's a draw!\n")
		}
	} else {
		if m.Current == board.PlayerX {
			out.WriteString("Player X's turn\n")
		} else {
			out.WriteString("Player O's turn\n")
		}
	}

	if m.ShowError {
		out.WriteString("\nCell already taken!\n")
	}

	block := lipgloss.Place(
		m.Width, m.Height,
		lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			out.String(),
			m.help.View(),
		),
		lipgloss.WithWhitespaceChars("@"),
		lipgloss.WithWhitespaceForeground(lipgloss.Color("#303033")),
	)

	return zone.Scan(block)
}

func New() Model {
	var player board.Cell
	start := rand.Int() % 2
	if start == 0 {
		player = board.PlayerX
	} else {
		player = board.PlayerO
	}
	return Model{
		Board:   board.NewBoard(),
		Current: player,
		help:    help.New(),
	}
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
