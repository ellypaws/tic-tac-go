package help

import (
	"github.com/charmbracelet/bubbles/key"
)

// The keyMap defines a set of keybindings. To work for help, it must satisfy
// key.Map. It could also very easily be a map[string]key.Binding.
type keyMap struct {
	Up         key.Binding
	Down       key.Binding
	Left       key.Binding
	Right      key.Binding
	Help       key.Binding
	Enter      key.Binding
	Reset      key.Binding
	AIMove     key.Binding
	ToggleAI   key.Binding
	Difficulty key.Binding
	Quit       key.Binding
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "move left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Reset: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "reset game"),
	),
	AIMove: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "Single AI move"),
	),
	ToggleAI: key.NewBinding(
		key.WithKeys("ctrl+a"),
		key.WithHelp("ctrl+a", "toggle AI"),
	),
	Difficulty: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "change difficulty"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
}
