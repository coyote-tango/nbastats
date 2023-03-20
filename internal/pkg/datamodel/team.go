package datamodel

import (
	"fmt"
	"strconv"
)

type Team struct {
	TeamId             int    `json:"teamId,omitempty"`
	Abbreviation       string `json:"abbreviation,omitempty"`
	City               string `json:"city,omitempty"`
	Nickname           string `json:"nickname,omitempty"`
	YearFounded        int    `json:"yearFounded,omitempty"`
	Arena              string `json:"arena,omitempty"`
	Owner              string `json:"owner,omitempty"`
	GeneralManager     string `json:"generalManager,omitempty"`
	HeadCoach          string `json:"headCoach,omitempty"`
	DLeagueAffiliation string `json:"dLeagueAffiliation,omitempty"`
}

func NewTeam(rowSet []interface{}) *Team {
	return &Team{
		TeamId:             int(rowSet[0].(float64)),
		Abbreviation:       rowSet[1].(string),
		Nickname:           rowSet[2].(string),
		YearFounded:        int(rowSet[3].(float64)),
		City:               rowSet[4].(string),
		Arena:              rowSet[5].(string),
		Owner:              rowSet[7].(string),
		GeneralManager:     rowSet[8].(string),
		DLeagueAffiliation: rowSet[10].(string),
	}
}

func GetTeam(id int) (*Team, error) {
	// https://stats.nba.com/stats/teamdetails
	resource, err := reqResource("teamdetails", map[string]string{
		"TeamId": strconv.Itoa(id)},
	)
	if err != nil {
		return nil, fmt.Errorf("error fetching team details: %w", err)
	}

	for _, resultSet := range resource.ResultSets {
		if resultSet.Name == "TeamBackground" {
			return NewTeam(resultSet.RowSet[0]), nil
		}
	}
	return nil, nil
}

func GetTeams() (*[]Team, error) {
	var teamsList []Team

	for abbreviation, teamInfo := range teams {
		teamsList = append(teamsList, Team{
			TeamId:       teamInfo.Id,
			Nickname:     teamInfo.Name,
			City:         teamInfo.City,
			Abbreviation: abbreviation,
		})
	}

	return &teamsList, nil
}
