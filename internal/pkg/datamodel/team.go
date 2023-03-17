package datamodel

import (
	"fmt"
	"strconv"
)

type Team struct {
	TeamId             int
	Abbreviation       string
	Nickname           string
	YearFounded        int
	City               string
	Arena              string
	Owner              string
	GeneralManager     string
	HeadCoach          string
	DLeagueAffiliation string
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

// Get team by team Id

func GetTeam(id int) (*Team, error) {
	// https://stats.nba.com/stats/teamdetails
	resource, err := getResource("teamdetails", map[string]string{"TeamId": strconv.Itoa(id)})
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
	teamsCh := make(chan Team)

	for _, value := range teams {
		go func(id int) {
			team, _ := GetTeam(id)
			fmt.Println(*team)
			teamsCh <- *team
		}(value.Id)
	}

	for range teams {
		teamsList = append(teamsList, <-teamsCh)
	}

	return &teamsList, nil
}
