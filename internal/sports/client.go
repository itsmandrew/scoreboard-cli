package sports

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL = "https://api.balldontlie.io/v1/games"

var gamesCache = NewCache[[]Game](30 * time.Second)

type GameResponse struct {
	Data []Game `json:"data"`
}

type Game struct {
	ID               int    `json:"id"`
	Status           string `json:"status"`
	Time             string `json:"time"`
	HomeTeamScore    int    `json:"home_team_score"`
	VisitorTeamScore int    `json:"visitor_team_score"`
	HomeTeam         Team   `json:"home_team"`
	VisitorTeam      Team   `json:"visitor_team"`
}

type Team struct {
	Abbreviation string `json:"abbreviation"`
}

func FetchNBAScores(apiKey string) ([]Game, error) {
	// Check cache first
	if cached, ok := gamesCache.Get(); ok {
		return cached, nil
	}

	// Cache miss or expired, fetch from API
	today := time.Now().Format("2006-01-02")

	url := fmt.Sprintf("%s?dates[]=%s", baseURL, today)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", apiKey)

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var gameData GameResponse

	if err := json.Unmarshal(body, &gameData); err != nil {
		return nil, err
	}

	// Update cache with fresh data
	gamesCache.Set(gameData.Data)

	return gameData.Data, nil
}

func FormatStatus(status string) string {
	t, err := time.Parse(time.RFC3339, status)
	if err != nil {
		return status
	}

	return t.Local().Format("3:04 PM")
}
