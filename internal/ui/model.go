package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/itsmandrew/scoreboard-tui/internal/sports"
)

type sessionState int

const (
	menuView sessionState = iota
	resultView
)

type LeagueLogo struct {
	Name  string
	Lines []string // 2 lines of ASCII art
	Color lipgloss.Color
}

var leagueLogos = []LeagueLogo{
	{
		Name: "NBA Games",
		Lines: []string{
			" ● ● ",
			" ███ ",
		},
		Color: NBAColor,
	},
	{
		Name: "NFL Games",
		Lines: []string{
			" ╔═╗ ",
			" ║ ║ ",
		},
		Color: NFLColor,
	},
	{
		Name: "NCAA Basketball",
		Lines: []string{
			" ╠══ ",
			" ╚══ ",
		},
		Color: NCAAColor,
	},
	{
		Name: "Exit",
		Lines: []string{
			" ╳ ╳ ",
			"  ╳  ",
		},
		Color: ExitColor,
	},
}

// Model holds the application state for the scoreboard TUI
type Model struct {
	state      sessionState  // Current view state (menu or result)
	cursor     int           // Menu cursor position
	selected   string        // Currently selected league
	loading    bool          // Loading state indicator
	spinner    spinner.Model // Loading spinner component
	apiKey     string        // API key for sports data
	nbaGames   []sports.Game // Fetched NBA game data
	gamesTable table.Model   // Scrollable table for displaying games
	err        error         // Error state
}

func InitialModel(apiKey string) Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(MainColor)

	return Model{
		state:   menuView,
		spinner: s,
		apiKey:  apiKey,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
