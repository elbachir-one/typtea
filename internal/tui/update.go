package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Update processes incoming messages and updates the model accordingly
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle window size changes
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	// Handle keyboard input and game logic
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "enter":
			if m.showResults {
				m.restartTest()
				return m, tickCmd()
			}
			return m, nil

		case " ":
			if !m.showResults && !m.game.IsFinished && !m.game.IsTimeUp() {
				m.game.AddCharacter(' ')
			}
			return m, nil

		case "backspace":
			if !m.showResults && !m.game.IsFinished {
				m.game.RemoveCharacter()
			}
			return m, nil

		default:
			// Handle regular character input
			if !m.showResults && !m.game.IsFinished && !m.game.IsTimeUp() {
				runes := []rune(msg.String())
				if len(runes) == 1 && runes[0] >= 32 && runes[0] <= 126 {
					m.game.AddCharacter(runes[0])
				}
			}
			return m, nil
		}

	// Handle tick messages for periodic updates
	case tickMsg:
		if !m.showResults {
			if m.game.IsTimeUp() && m.game.IsStarted {
				m.finalStats = m.game.GetStats()
				m.showResults = true
				return m, nil
			}
			return m, tickCmd()
		}
		return m, nil
	}

	return m, nil
}
