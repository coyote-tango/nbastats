package datamodel

import (
	"fmt"
	"time"
)

// https://cdn.nba.com/static/json/liveData/odds/odds_todaysGames.json
// https://stats.nba.com/stats/boxscoresummaryv2?GameID=0022200819

type Game struct {
	GameId      string    `json:"gameId"`
	GameTimeUTC time.Time `json:"gameTimeUTC"`
	GameTimeCT  time.Time `json:"gameTimeCT"`
	HomeTeam    Team      `json:"homeTeam"`
	AwayTeam    Team      `json:"awayTeam"`
}

func GetGames(gameDate time.Time) (*[]Game, error) {
	//https://stats.nba.com/stats/scoreboardv3?GameDate=2023-04-01&LeagueID=00
	var params = map[string]string{"GameDate": gameDate.Format("2006-01-02")}
	gamesResource, err := reqGamesResource(params)
	if err != nil {
		return nil, fmt.Errorf("error fetching game details: %w", err)
	}

	gamesSlice := gamesResource.([]interface{})
	var gamesList []Game
	for _, game := range gamesSlice {
		gameMap := game.(map[string]interface{})
		homeTeamMap := gameMap["homeTeam"].(map[string]interface{})
		awayTeamMap := gameMap["awayTeam"].(map[string]interface{})
		gameTimeUTC, _ := time.Parse("2006-01-02T15:04:05Z07:00", gameMap["gameTimeUTC"].(string))
		location, _ := time.LoadLocation("America/Mexico_City")
		gameTimeCT := gameTimeUTC.In(location)

		gamesList = append(gamesList, Game{
			GameId:      gameMap["gameId"].(string),
			GameTimeUTC: gameTimeUTC,
			GameTimeCT:  gameTimeCT,
			HomeTeam: Team{
				TeamId:       int(homeTeamMap["teamId"].(float64)),
				Abbreviation: homeTeamMap["teamTricode"].(string),
				City:         homeTeamMap["teamCity"].(string),
				Nickname:     homeTeamMap["teamName"].(string),
			},
			AwayTeam: Team{
				TeamId:       int(homeTeamMap["teamId"].(float64)),
				Abbreviation: awayTeamMap["teamTricode"].(string),
				City:         awayTeamMap["teamCity"].(string),
				Nickname:     awayTeamMap["teamName"].(string),
			},
		})
	}
	print(gamesResource)
	return &gamesList, nil
}
