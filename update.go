package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"
	"github.com/lrstanley/bubblezone"
	"tic-tac-toe/ai"
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
		keys := m.help.Keys
		switch {
		case key.Matches(msg, keys.Reset):
			newModel := New(
				WithAI(m.ai),
				WithDifficulty(m.ai.Difficulty),
				WithStartPlayer(m.Current),
				WithHeight(m.Height),
				WithWidth(m.Width),
			)
			return &newModel, nil

			//m.Board = board.NewBoard()
			//m.gameOver = false
			//return m, nil
		case key.Matches(msg, keys.Up):
			m.Cursor.Col = max(m.Cursor.Col-1, 0)
		case key.Matches(msg, keys.Down):
			m.Cursor.Col = min(m.Cursor.Col+1, 2)
		case key.Matches(msg, keys.Left):
			m.Cursor.Row = max(m.Cursor.Row-1, 0)
		case key.Matches(msg, keys.Right):
			m.Cursor.Row = min(m.Cursor.Row+1, 2)
		case key.Matches(msg, keys.Enter):
			if m.checkWinner() {
				return m, nil
			}
			if m.Board.Cells[m.Cursor.Col][m.Cursor.Row] == board.Empty {
				m.Board.Cells[m.Cursor.Col][m.Cursor.Row] = m.Current
				m.swapPlayer()
			} else {
				m.invalidMove = true
				return m, nil
			}
			if m.checkWinner() {
				return m, nil
			}
		case key.Matches(msg, keys.Difficulty):
			m.ai.ChangeDifficulty(ai.Difficulty((int(m.ai.Difficulty) + 1) % 5))
		case key.Matches(msg, keys.AIMove):
			if m.checkWinner() {
				return m, nil
			}
			x, y := m.ai.Move(m.Board)
			m.Board.Cells[x][y] = m.Current
			m.swapPlayer()
			if m.checkWinner() {
				return m, nil
			}
		}
		m.help, _ = m.help.Update(msg)
		return m, nil
	case tea.MouseMsg:
		done := m.mouseUpdate(msg)
		if done {
			return m, nil
		}
	}
	return m, nil
}

func (m *Model) mouseUpdate(msg tea.MouseMsg) bool {
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
			return true
		}
		if m.Board.Cells[m.Cursor.Col][m.Cursor.Row] == board.Empty {
			m.Board.Cells[m.Cursor.Col][m.Cursor.Row] = m.Current
			m.swapPlayer()
		} else {
			m.invalidMove = true
			return true
		}
		if m.checkWinner() {
			return true
		}
	}
	return false
}

func (m *Model) swapPlayer() {
	if m.Current == board.PlayerX {
		m.Current = board.PlayerO
	} else {
		m.Current = board.PlayerX
	}
}

func (m *Model) checkWinner() bool {
	winner := m.Board.Winner()
	if winner != board.Empty || m.Board.IsFull() {
		m.gameOver = true
		return true
	}
	return false
}
