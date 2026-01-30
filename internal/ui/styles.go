package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	MainColor   = lipgloss.Color("#7D56F4")
	SubtleColor = lipgloss.Color("241")
	ErrorColor  = lipgloss.Color("#FF5555")
	White       = lipgloss.Color("#FFFDF5")

	// Highlight color
	HighlightColor = lipgloss.Color("#D4AF87") // Warm beige for selections

	// League-specific colors
	NBAColor  = lipgloss.Color("#FF6B35") // Orange
	NFLColor  = lipgloss.Color("#4ECDC4") // Teal
	NCAAColor = lipgloss.Color("#5DADE2") // Blue
	ExitColor = lipgloss.Color("#95A5A6") // Gray

	// Base Containers
	WindowStyle = lipgloss.NewStyle().
			Border(lipgloss.ThickBorder()).
			BorderForeground(SubtleColor).
			Padding(1, 4).
			Margin(2)

	// Text Styles
	TitleStyle = lipgloss.NewStyle().
			Foreground(White).
			Border(lipgloss.DoubleBorder()).
			BorderForeground(White).
			Padding(1, 3).
			Bold(true).
			Align(lipgloss.Center).
			Width(60)

	ErrorTitleStyle = lipgloss.NewStyle().
				Foreground(White).
				Background(ErrorColor).
				Padding(1, 0).
				Bold(true).
				Align(lipgloss.Center).
				Width(60)

	// Added missing subtle and italic styles
	SubtleStyle = lipgloss.NewStyle().
			Foreground(SubtleColor)

	ItalicStyle = lipgloss.NewStyle().
			Italic(true).
			Foreground(SubtleColor)

	// Logo box styles
	LogoBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(SubtleColor).
			Padding(0, 1).
			Width(7).
			Height(2).
			Align(lipgloss.Center)

	LogoBoxSelectedStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(HighlightColor).
				Padding(0, 1).
				Width(7).
				Height(2).
				Align(lipgloss.Center).
				Bold(true)

	// Menu label styles
	MenuLabelStyle = lipgloss.NewStyle().
			Width(12).
			Align(lipgloss.Center)

	MenuLabelSelectedStyle = lipgloss.NewStyle().
				Width(12).
				Align(lipgloss.Center).
				Foreground(HighlightColor).
				Bold(true)
)
