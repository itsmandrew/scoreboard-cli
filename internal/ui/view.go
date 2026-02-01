package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"github.com/itsmandrew/scoreboard-tui/internal/sports"
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
			// Keep original logo color when selected
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

	// Center the menu within the container
	centeredMenu := lipgloss.NewStyle().
		Width(60).
		Align(lipgloss.Center).
		Render(menu)

	b.WriteString(centeredMenu + "\n\n")

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
	var title string
	var content string

	if m.selected == "NBA Games" {
		title = TitleStyle.Render(fmt.Sprintf(" NBA Scores - %s ", time.Now().Format("Jan 02, 2006")))
		content = m.renderNBAGamesTable()
	} else {
		title = TitleStyle.Render(" SCOREBOARD ")
		content = fmt.Sprintf("Returning data for: %s", m.selected)
	}

	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		title,
		content,
		ItalicStyle.Render("(↑/↓ to scroll • Enter to go back)"),
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

// renderNBAGamesTable displays NBA games in a scrollable table format
func (m Model) renderNBAGamesTable() string {
	if len(m.nbaGames) == 0 {
		return SubtleStyle.Render("No games scheduled for today")
	}

	return m.gamesTable.View()
}

// createNBATable builds a scrollable table from NBA game data with custom styling
func createNBATable(games []sports.Game) table.Model {
	// Define table columns
	columns := []table.Column{
		{Title: "Home", Width: 8},
		{Title: "Away", Width: 8},
		{Title: "Score", Width: 14},
		{Title: "Status", Width: 14},
		{Title: "Time", Width: 14},
	}

	// Build rows from game data
	var rows []table.Row
	for _, game := range games {
		rows = append(rows, table.Row{
			game.HomeTeam.Abbreviation,
			game.VisitorTeam.Abbreviation,
			fmt.Sprintf("(%d - %d)", game.HomeTeamScore, game.VisitorTeamScore),
			sports.FormatStatus(game.Status),
			game.Time,
		})
	}

	// Create table with focused state and fixed height
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	// Apply custom styling for beige row highlighting
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("241")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("16")). // Black text for contrast
		Background(HighlightColor).       // Beige background for selected row
		Bold(true)
	t.SetStyles(s)

	return t
}
