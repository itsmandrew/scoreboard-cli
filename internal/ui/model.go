package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/itsmandrew/scoreboard-cli/internal/sports"
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

type Model struct {
	state    sessionState
	choices  []string
	cursor   int
	selected string
	loading  bool
	spinner  spinner.Model
	apiKey   string
	nbaGames []sports.Game
	err      error
}

func InitialModel(apiKey string) Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(MainColor)

	return Model{
		state:   menuView,
		choices: []string{"NBA Games", "NFL Games", "NCAA Basketball", "Exit"},
		spinner: s,
		apiKey:  apiKey,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
