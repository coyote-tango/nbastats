package datamodel

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type Player struct {
	PlayerId     int        `json:"playerId,omitempty"`
	FirstName    string     `json:"firstName,omitempty"`
	LastName     string     `json:"lastName,omitempty"`
	DisplayName  string     `json:"displayName,omitempty"`
	TeamId       int        `json:"teamId,omitempty"`
	TeamNickname string     `json:"teamNickname,omitempty"`
	BirthDate    *time.Time `json:"birthDate,omitempty"`
	Country      string     `json:"country,omitempty"`
	Height       float64    `json:"height,omitempty"`
	Weight       float64    `json:"weight,omitempty"`
	Jersey       string     `json:"jersey,omitempty"`
	Position     string     `json:"position,omitempty"`
	DraftYear    int        `json:"draftYear,omitempty"`
	DraftRound   int        `json:"draftRound,omitempty"`
	DraftNumber  int        `json:"draftNumber,omitempty"`
}

func NewPlayer(rowSet []interface{}) *Player {
	birthDate, _ := time.Parse("2006-01-02T15:04:05", rowSet[7].(string))
	draftYear, _ := strconv.Atoi(rowSet[29].(string))
	draftRound, _ := strconv.Atoi(rowSet[30].(string))
	draftNumber, _ := strconv.Atoi(rowSet[31].(string))

	return &Player{
		PlayerId:    int(rowSet[0].(float64)),
		FirstName:   rowSet[1].(string),
		LastName:    rowSet[2].(string),
		DisplayName: rowSet[3].(string),
		BirthDate:   &birthDate,
		Country:     rowSet[9].(string),
		Height:      parseHeight(rowSet[11].(string)),
		Weight:      parseWeight(rowSet[12].(string)),
		Jersey:      rowSet[14].(string),
		Position:    rowSet[15].(string),
		TeamId:      int(rowSet[18].(float64)),
		DraftYear:   draftYear,
		DraftRound:  draftRound,
		DraftNumber: draftNumber,
	}
}

func GetPlayers() (*[]Player, error) {
	//https://stats.nba.com/stats/commonallplayers?IsOnlyCurrentSeason=1&LeagueID=00&Season=2022-23
	var playersList []Player
	resource, err := getResource("commonallplayers", map[string]string{"IsOnlyCurrentSeason": "1",
		"LeagueID": NBA, "Season": CURRENT})
	if err != nil {
		return nil, fmt.Errorf("error fetching player details: %w", err)
	}

	for _, resultSet := range resource.ResultSets {
		if resultSet.Name == "CommonAllPlayers" {
			for _, row := range resultSet.RowSet { // Iterate through rows, not just the first row
				playersList = append(playersList, Player{
					PlayerId:     int(row[0].(float64)),
					DisplayName:  row[2].(string),
					TeamNickname: row[10].(string),
					TeamId:       int(row[8].(float64)),
				})
			}
		}
	}
	return &playersList, nil

}

func GetPlayer(id int) (*Player, error) {
	//https://stats.nba.com/stats/commonplayerinfo?LeagueID=&PlayerID=2544
	resource, err := getResource("commonplayerinfo", map[string]string{"PlayerId": strconv.Itoa(id)})
	if err != nil {
		return nil, fmt.Errorf("error fetching player details: %w", err)
	}

	for _, resultSet := range resource.ResultSets {
		if resultSet.Name == "CommonPlayerInfo" {
			return NewPlayer(resultSet.RowSet[0]), nil
		}
	}
	return nil, nil

}

func parseHeight(heightStr string) float64 {
	parts := strings.Split(heightStr, "-")
	if len(parts) != 2 {
		return 0
	}

	feet, _ := strconv.Atoi(parts[0])
	inches, _ := strconv.Atoi(parts[1])

	totalInches := float64(feet*12 + inches)
	cm := totalInches * 2.54

	return math.Round(cm)
}

func parseWeight(weightStr string) float64 {
	pounds, _ := strconv.ParseFloat(weightStr, 64)
	kg := pounds * 0.45359237

	return math.Round(kg)
}
