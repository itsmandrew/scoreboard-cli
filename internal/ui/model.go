package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type sessionState int

const (
	menuView sessionState = iota
	resultView
)

type Model struct {
	state    sessionState
	choices  []string
	cursor   int
	selected string
}

func InitialModel() Model {
	return Model{
		state:   menuView,
		choices: []string{"NFL Games", "NBA Games", "NCAA Basketball", "Exit"},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.state == menuView && m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.state == menuView && m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			if m.state == menuView {
				m.selected = m.choices[m.cursor]
				if m.selected == "Exit" {
					return m, tea.Quit
				}
				m.state = resultView
			} else {
				m.state = menuView
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	var body string

	if m.state == resultView {
		// The "Message" Screen
		body = fmt.Sprintf(
			"%s\n\nReturning data for: %s\n\n%s",
			TitleStyle.Render(" SYSTEM MESSAGE "),
			m.selected,
			"Press Enter to go back...",
		)
	} else {
		// The "Menu" Screen
		var b strings.Builder
		b.WriteString(TitleStyle.Render(" TICKER MAIN MENU ") + "\n\n")

		for i, choice := range m.choices {
			cursor := " "
			label := choice

			if m.cursor == i {
				cursor = ">"
				label = lipgloss.NewStyle().Foreground(MainColor).Bold(true).Render(choice)
			}
			b.WriteString(fmt.Sprintf("%s %s\n", cursor, label))
		}

		b.WriteString("\n(j/k to move • enter to select)")
		body = b.String()
	}

	return WindowStyle.Render(body)
}
