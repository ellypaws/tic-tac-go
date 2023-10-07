package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"
	"github.com/lrstanley/bubblezone"
	"tic-tac-toe/board"
)

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
			if m.checkWinner() {
				return m, nil
			}
			if m.Board.Cells[m.Cursor.Col][m.Cursor.Row] == board.Empty {
				m.Board.Cells[m.Cursor.Col][m.Cursor.Row] = m.Current
				if m.Current == board.PlayerX {
					m.Current = board.PlayerO
				} else {
					m.Current = board.PlayerX
				}
			} else {
				m.invalidMove = true
				return m, nil
			}
			if m.checkWinner() {
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
		if msg.Type == tea.MouseLeft {
			if m.checkWinner() {
				return m, nil
			}
			if m.Board.Cells[m.Cursor.Col][m.Cursor.Row] == board.Empty {
				m.Board.Cells[m.Cursor.Col][m.Cursor.Row] = m.Current
				if m.Current == board.PlayerX {
					m.Current = board.PlayerO
				} else {
					m.Current = board.PlayerX
				}
			} else {
				m.invalidMove = true
				return m, nil
			}
			if m.checkWinner() {
				return m, nil
			}
		}
	}
	return m, nil
}

func (m *Model) checkWinner() bool {
	winner := m.Board.Winner()
	if winner != board.Empty || m.Board.IsFull() {
		m.gameOver = true
		return true
	}
	return false
}
