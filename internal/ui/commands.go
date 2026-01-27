package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/itsmandrew/scoreboard-cli/internal/sports"
)

type (
	nbaMsg []sports.Game
	errMsg error
)

func FetchNBACmd(apiKey string) tea.Cmd {
	return func() tea.Msg {
		games, err := sports.FetchNBAScores(apiKey)
		if err != nil {
			return errMsg(err)
		}

		return nbaMsg(games)
	}
}
