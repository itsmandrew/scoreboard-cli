package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
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
	loading  bool
	spinner  spinner.Model
}

func InitialModel() Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(MainColor)

	return Model{
		state:   menuView,
		choices: []string{"NFL Games", "NBA Games", "NCAA Basketball", "Exit"},
		spinner: s,
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
		}
		if m.loading {
			return m, nil
		}

		return m.handleKeyInput(msg)

	case string:
		m.loading = false
		m.state = resultView
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m *Model) handleKeyInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.state == resultView {
		if msg.String() == "enter" {
			m.state = menuView
		}
		return m, nil
	}

	// Menu Navigation Logic
	switch msg.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.choices)-1 {
			m.cursor++
		}
	case "enter":
		m.selected = m.choices[m.cursor]
		if m.selected == "Exit" {
			return m, tea.Quit
		}
		m.loading = true
		return m, tea.Batch(m.spinner.Tick, m.simulateFetch())
	}

	return m, nil
}

func (m Model) simulateFetch() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(3 * time.Second)
		return "Done"
	}
}

func (m Model) View() string {
	var body string
	if m.loading {
		body = fmt.Sprintf(
			"\n %s  Fetching %s data...\n\n%s",
			m.spinner.View(),
			m.selected,
			lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render("Please wait..."),
		)
	} else if m.state == resultView {
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
		b.WriteString(TitleStyle.Render(" SCOREBOARD CLI MENU ") + "\n\n")

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
