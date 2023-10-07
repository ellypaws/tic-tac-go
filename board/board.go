package board

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"strconv"
)

type Board struct {
	columns map[string]table.Model
	spinner spinner.Model
}

var boardSymbol = map[string]string{
	"X": "❌",
	"O": "⭕",
}

var focused = table.DefaultStyles()

func NewCol(prefix int) table.Model {
	columns := []table.Column{
		{Title: "[ ]", Width: 6},
	}

	rows := []table.Row{
		{"AGH"},
		{"AGH"},
		{"AGH"},
	}

	for row, _ := range rows {
		for col, cell := range rows[row] {
			rows[row][col] = zone.Mark(strconv.Itoa(prefix)+strconv.Itoa(row)+strconv.Itoa(col), cell)
		}
	}

	newTable := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(len(rows)+2),
	)

	return newTable
}

func NewBoard() Board {
	focused = table.DefaultStyles()
	focused.Header = focused.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	focused.Selected = focused.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	newTable := make(map[string]table.Model)

	for i := 0; i < 3; i++ {
		newTable[strconv.Itoa(i)] = NewCol(i)
	}

	s := spinner.New()
	s.Spinner = spinner.Dot

	return Board{
		columns: newTable,
		spinner: s,
	}
}

func (b Board) View() string {
	var view string
	for _, col := range b.columns {
		view += col.View()
	}
	return view
}

func (b Board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		msg.Width -= 4
		msg.Height -= 2
		return b, nil
	case tea.MouseMsg:
		//if msg.Type == tea.MouseLeft {
		//	for i := 0; i < 3; i++ {
		//		for j := 0; j < 3; j++ {
		//			if zone.Get(b.columns[i*3+j].Id).InBounds(msg) {
		//				newCol, _ := b.columns[i*3+j].Update(msg)
		//				b.columns[i*3+j] = newCol.(table.Model)
		//				return b, nil
		//			}
		//		}
		//	}
		//}
	}
	for i, k := range b.columns {
		newCol, _ := k.Update(msg)
		b.columns[i] = newCol
	}
	return b, cmd
}

func (b Board) Init() tea.Cmd {
	return nil
}
