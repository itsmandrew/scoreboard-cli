package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/itsmandrew/scoreboard-cli/internal/ui"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, proceeding w/ systen env")
	}

	apiKey := os.Getenv("BALLDONTLIE_API_KEY")

	if apiKey == "" {
		log.Fatal("API KEY is not set")
	}

	p := tea.NewProgram(ui.InitialModel(apiKey))
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's an error: %v", err)
		os.Exit(1)
	}
}
