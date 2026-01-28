package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/itsmandrew/scoreboard-cli/internal/sports"
)

func (m Model) View() string {
	var body string

	// Route rendering based on application state
	switch {
	case m.err != nil:
		body = m.errorView()
	case m.loading:
		body = m.loadingView()
	case m.state == resultView:
		body = m.resultView()
	default:
		body = m.menuView()
	}

	return WindowStyle.Render(body)
}

func (m Model) menuView() string {
	var b strings.Builder
	b.WriteString(TitleStyle.Render(" ⚡ SCOREBOARD ⚡ ") + "\n\n")

	// Create all cards
	var cards []string
	for i, logo := range leagueLogos {
		// Choose styles based on selection
		boxStyle := LogoBoxStyle
		labelStyle := MenuLabelStyle
		logoColor := logo.Color

		if m.cursor == i {
			boxStyle = LogoBoxSelectedStyle
			labelStyle = MenuLabelSelectedStyle
			logoColor = MainColor // Purple when selected
		}

		// Build logo box content with colored symbols
		var logoLines []string
		for _, line := range logo.Lines {
			logoLines = append(logoLines,
				lipgloss.NewStyle().Foreground(logoColor).Render(line))
		}
		logoContent := strings.Join(logoLines, "\n")

		// Render the logo box and label
		logoBox := boxStyle.Render(logoContent)

		// Simplify label (just league abbreviation)
		labelText := logo.Name
		if logo.Name == "NBA Games" {
			labelText = "NBA"
		} else if logo.Name == "NFL Games" {
			labelText = "NFL"
		} else if logo.Name == "NCAA Basketball" {
			labelText = "NCAA"
		} else {
			labelText = "Exit"
		}
		label := labelStyle.Render(labelText)

		// Create cursor indicator below selected item
		cursor := labelStyle.Render(" ")
		if m.cursor == i {
			cursor = labelStyle.Render("^")
		}

		// Stack logo box, label, and cursor vertically
		card := lipgloss.JoinVertical(
			lipgloss.Center,
			logoBox,
			label,
			cursor,
		)

		cards = append(cards, card)
	}

	// Join all cards horizontally with spacing
	menu := lipgloss.JoinHorizontal(lipgloss.Top, cards...)
	b.WriteString(menu + "\n\n")

	b.WriteString(fmt.Sprintf("%s", SubtleStyle.Render("(h/l or ←/→ to move • enter to select)")))
	return b.String()
}

func (m Model) loadingView() string {
	return fmt.Sprintf(
		"\n %s  Fetching %s data...\n\n%s",
		m.spinner.View(),
		m.selected,
		SubtleStyle.Render("Please wait..."),
	)
}

func (m Model) resultView() string {
	content := ""
	if m.selected == "NBA Games" {
		content = m.renderNBAGames()
	} else {
		content = fmt.Sprintf("Returning data for: %s", m.selected)
	}

	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		TitleStyle.Render(" SCOREBOARD "),
		content,
		ItalicStyle.Render("Press Enter to go back..."),
	)
}

func (m Model) errorView() string {
	return fmt.Sprintf(
		"%s\n\nError: %v\n\n%s",
		ErrorTitleStyle.Render(" ERROR "),
		m.err,
		ItalicStyle.Render("Press Enter to go back..."),
	)
}

// Format fetched NBA game data into a table-like string
func (m Model) renderNBAGames() string {
	if len(m.nbaGames) == 0 {
		return SubtleStyle.Render("No games scheduled for today")
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("NBA Scores - %s\n\n", time.Now().Format("Jan 02, 2006")))

	for _, game := range m.nbaGames {
		home := TeamStyle.Render(fmt.Sprintf("%s %d", game.HomeTeam.Abbreviation, game.HomeTeamScore))
		away := TeamStyle.Render(fmt.Sprintf("%s %d", game.VisitorTeam.Abbreviation, game.VisitorTeamScore))
		status := StatusStyle.Render(sports.FormatStatus(game.Status))

		b.WriteString(fmt.Sprintf("%s vs %s — %s\n", home, away, status))
	}
	return b.String()
}
