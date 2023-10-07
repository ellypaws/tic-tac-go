package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/lrstanley/bubblezone"
	"strings"
	"tic-tac-toe/ai"
	"tic-tac-toe/board"
)

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
		Padding(0, 0)

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

	if m.gameOver {
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

	if m.invalidMove {
		out.WriteString("\nCell already taken!\n")
	} else {
		out.WriteString("\n\n")
	}

	var difficulty strings.Builder

	switch m.ai.Difficulty {
	case ai.Easy:
		difficulty.WriteString("Easy")
	case ai.Medium:
		difficulty.WriteString("Medium")
	case ai.Hard:
		difficulty.WriteString("Hard")
	case ai.Minimax:
		difficulty.WriteString("Minimax")
	case ai.Neural:
		difficulty.WriteString("Neural Network")
	}

	block := lipgloss.Place(
		m.Width, m.Height,
		lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			difficulty.String(),
			out.String(),
			m.help.View(),
		),
		lipgloss.WithWhitespaceChars("@"),
		lipgloss.WithWhitespaceForeground(lipgloss.Color("#303033")),
	)

	return zone.Scan(block)
}
