package board

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"strconv"
)

// Constants representing player symbols and an empty cell
const (
	EmptySymbol = " "
	XSymbol     = "❌"
	OSymbol     = "⭕"
)

// Cell represents each cell of the Tic Tac Toe board.
type Cell struct {
	Value string
	ID    string
}

// Board represents the Tic Tac Toe game board.
type Board struct {
	Cells [3][3]Cell
}

// NewBoard initializes a new game board.
func NewBoard() Board {
	b := Board{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			id := strconv.Itoa(i) + "-" + strconv.Itoa(j)
			b.Cells[i][j] = Cell{Value: EmptySymbol, ID: zone.Mark(id, "   ")}
		}
	}

	return b
}

// View renders the game board as a string.
func (b Board) View() string {
	var view string

	cellStyle := lipgloss.NewStyle().Width(5).Height(3).Border(lipgloss.NormalBorder())
	dividerStyle := lipgloss.NewStyle().Width(17)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			cell := b.Cells[i][j]
			view += cellStyle.Render(cell.Value)

			if j < 2 {
				view += " "
			}
		}
		if i < 2 {
			view += "\n" + dividerStyle.Render("-+-+-") + "\n"
		}
	}

	return view
}

// Update handles mouse input to update the board state.
func (b Board) Update(msg tea.Msg) (Board, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseMsg:
		if msg.Type == tea.MouseLeft {
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					cell := &b.Cells[i][j]

					if zone.Get(cell.ID).InBounds(msg) {
						switch cell.Value {
						case EmptySymbol:
							cell.Value = XSymbol
						case XSymbol:
							cell.Value = OSymbol
						default:
							cell.Value = EmptySymbol
						}
						return b, nil
					}
				}
			}
		}
	}
	return b, nil
}

// Init initializes the board model; currently a no-op.
func (b Board) Init() tea.Cmd {
	return nil
}
