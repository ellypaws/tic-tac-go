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
			if m.Cursor.Y > 0 {
				m.Cursor.Y--
			}
		case tea.KeyDown:
			if m.Cursor.Y < 2 {
				m.Cursor.Y++
			}
		case tea.KeyLeft:
			if m.Cursor.X > 0 {
				m.Cursor.X--
			}
		case tea.KeyRight:
			if m.Cursor.X < 2 {
				m.Cursor.X++
			}
		case tea.KeyEnter:
			if m.Board.Cells[m.Cursor.Y][m.Cursor.X] == Empty {
				m.Board.Cells[m.Cursor.Y][m.Cursor.X] = m.Current
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

	cursorStyle := lipgloss.NewStyle().Background(lipgloss.Color("#C44358"))

	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			zoneID := fmt.Sprintf("%d-%d", row, col)

			var cell string
			switch m.Board.Cells[row][col] {
			case PlayerX:
				cell = "X"
			case PlayerO:
				cell = "O"
			default:
				cell = " "
			}

			if row == m.Cursor.Y && col == m.Cursor.X {
				cell = cursorStyle.Render(cell)
			}

			out.WriteString(zone.Mark(m.Board.Id+zoneID, cell))

			if col < 2 {
				out.WriteString("|")
			}
		}
		if row < 2 {
			out.WriteString("\n-+-+-\n")
		}
	}

	out.WriteString("\n\n")

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
