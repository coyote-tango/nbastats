package datamodel

import "time"

// https://cdn.nba.com/static/json/liveData/odds/odds_todaysGames.json
// https://stats.nba.com/stats/boxscoresummaryv2?GameID=0022200819

type Game struct {
	GameId     int64
	DateEst    time.Time
	HomeTeamId int64
	AwayTeamId int64
	Season     string
	Attendance string
}
