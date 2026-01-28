package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		return m.handleKeyEvents(msg)

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case nbaMsg:
		m.loading = false
		m.state = resultView
		m.nbaGames = msg
		return m, nil

	case errMsg:
		m.loading = false
		m.err = msg
		return m, nil

	case string: // For simulated "DONE" messages
		m.loading = false
		m.state = resultView
		return m, nil
	}

	return m, nil
}

func (m Model) handleKeyEvents(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if msg.String() == "ctrl+c" || msg.String() == "q" {
		return m, tea.Quit
	}

	if m.loading {
		return m, nil
	}

	if m.state == resultView {
		if msg.String() == "enter" {
			m.state = menuView
			m.err = nil
		}
		return m, nil
	}

	return m.handleMenuNavigation(msg)
}

func (m Model) handleMenuNavigation(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "left", "h":
		if m.cursor > 0 {
			m.cursor--
		}
	case "right", "l":
		if m.cursor < len(leagueLogos)-1 {
			m.cursor++
		}
	case "enter":
		m.selected = leagueLogos[m.cursor].Name
		if m.selected == "Exit" {
			return m, tea.Quit
		}

		m.loading = true
		if m.selected == "NBA Games" {
			return m, tea.Batch(m.spinner.Tick, FetchNBACmd(m.apiKey))
		}
		return m, tea.Batch(m.spinner.Tick, m.simulateFetch())
	}
	return m, nil
}

func (m Model) simulateFetch() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(2 * time.Second)
		return "Done"
	}
}
