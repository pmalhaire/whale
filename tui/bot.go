package tui

import (
	"whale/game"

	tea "github.com/charmbracelet/bubbletea"
)

// Bot plays an action for a bot
func (m model) Bot(msg tea.Msg, player *game.Player) bool {

	// TODO fix coupling with m.actions
	// always plays the first available action
	// TODO fix bot new play bonuses
	player.Play(m.game.Deck, m.actions[0], nil, nil)
	return true
}
