package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lrstanley/bubblezone"
)

const (
	Empty = iota
	PlayerX
	PlayerO
)

type Cell int

type Board struct {
	Cells [3][3]Cell
	Id    string
}

func NewBoard() *Board {
	return &Board{
		Id: zone.NewPrefix(),
	}
}

func (b *Board) IsFull() bool {
	for _, row := range b.Cells {
		for _, cell := range row {
			if cell == Empty {
				return false
			}
		}
	}
	return true
}

func (b *Board) Winner() Cell {
	for i := 0; i < 3; i++ {
		if b.Cells[i][0] != Empty && b.Cells[i][0] == b.Cells[i][1] && b.Cells[i][0] == b.Cells[i][2] {
			return b.Cells[i][0]
		}
		if b.Cells[0][i] != Empty && b.Cells[0][i] == b.Cells[1][i] && b.Cells[0][i] == b.Cells[2][i] {
			return b.Cells[0][i]
		}
	}
	if b.Cells[0][0] != Empty && b.Cells[0][0] == b.Cells[1][1] && b.Cells[0][0] == b.Cells[2][2] {
		return b.Cells[0][0]
	}
	if b.Cells[0][2] != Empty && b.Cells[0][2] == b.Cells[1][1] && b.Cells[0][2] == b.Cells[2][0] {
		return b.Cells[0][2]
	}
	return Empty
}

type Model struct {
	Board         *Board
	Current       Cell
	GameOver      bool
	ShowError     bool
	Cursor        struct{ Row, Col int }
	Width, Height int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
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
			if m.Board.Cells[m.Cursor.Col][m.Cursor.Row] == Empty {
				m.Board.Cells[m.Cursor.Col][m.Cursor.Row] = m.Current
				if m.Current == PlayerX {
					m.Current = PlayerO
				} else {
					m.Current = PlayerX
				}
			} else {
				m.ShowError = true
				return m, nil
			}
		}

		winner := m.Board.Winner()
		if winner != Empty || m.Board.IsFull() {
			m.GameOver = true
		}

		return m, nil
	case tea.MouseMsg:
		if msg.Type == tea.MouseLeft {
			for row := 0; row < 3; row++ {
				for col := 0; col < 3; col++ {
					cell := &m.Board.Cells[row][col]
					if zone.Get(m.Board.Id + fmt.Sprintf("%d-%d", row, col)).InBounds(msg) {
						if *cell == Empty {
							*cell = m.Current
							if m.Current == PlayerX {
								m.Current = PlayerO
							} else {
								m.Current = PlayerX
							}
						} else {
							m.ShowError = true
							return m, nil
						}
					}
				}
			}
		}
	}
	return m, nil
}

//const (
//	EmptySymbol = "   "
//	XSymbol     = " ❌ "
//	OSymbol     = " ⭕ "
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
		Foreground(lipgloss.Color("#ff0000"))

	cursorStyle := normalStyle.Copy().
		BorderForeground(lipgloss.Color("86"))

	out.WriteString(fmt.Sprintf("[%v]:[%v] {%v:%v}\n", m.Cursor.Row, m.Cursor.Col, m.Width, m.Height))

	for col := 0; col < 3; col++ {
		rowItems := []string{}
		for row := 0; row < 3; row++ {
			//zoneID := fmt.Sprintf("%d-%d", row, col)

			var r = normalStyle.Render
			if col == m.Cursor.Col && row == m.Cursor.Row {
				r = cursorStyle.Render
			}

			var cell = r(EmptySymbol)
			//
			switch m.Board.Cells[col][row] {
			case PlayerX:
				cell = r(XSymbol)
			case PlayerO:
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
		case PlayerX:
			out.WriteString("Player X wins!\n")
		case PlayerO:
			out.WriteString("Player O wins!\n")
		default:
			out.WriteString("It's a draw!\n")
		}
	} else {
		if m.Current == PlayerX {
			out.WriteString("Player X's turn\n")
		} else {
			finalOut.WriteString("Player O's turn\n")
		}
	}

	if m.ShowError {
		finalOut.WriteString("\nCell already taken!\n")
	}

	block := lipgloss.Place(
		m.Width, m.Height,
		lipgloss.Center, lipgloss.Center,
		out.String(),
		lipgloss.WithWhitespaceChars("    ⭒"),
		lipgloss.WithWhitespaceForeground(lipgloss.Color("#757575")),
	)
	//block = lipgloss.PlaceVertical(m.Height, lipgloss.Center, block)
	out.WriteString(block)
	//out.WriteString("\n\n")

	return finalOut.String()
}

func main() {
	zone.NewGlobal()
	m := Model{
		Board:   NewBoard(),
		Current: PlayerX,
	}
	p := tea.NewProgram(
		&m,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if err, _ := p.Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
