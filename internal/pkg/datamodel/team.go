package datamodel

import (
	"context"
)

type Team struct {
	TeamId             float64
	Abbreviation       string
	Nickname           string
	YearFounded        float64
	City               string
	Arena              string
	Owner              string
	GeneralManager     string
	HeadCoach          string
	DLeagueAffiliation string
}

var teams []Team

func GetTeam(ctx context.Context, id string) (*Team, error) {
	// https://stats.nba.com/stats/teamdetails

	resource, _ := getResource("teamdetails", map[string]string{"TeamId": id})

	team := &Team{}

	for _, resultSet := range resource.ResultSets {
		if resultSet.Name == "TeamBackground" {
			for _, row := range resultSet.RowSet {
				team.TeamId, _ = row[0].(float64)
				team.Abbreviation = row[1].(string)
				team.Nickname = row[2].(string)
				team.YearFounded = row[3].(float64)
				team.City = row[4].(string)
				team.Arena = row[5].(string)
				team.Owner = row[7].(string)
				team.GeneralManager = row[8].(string)
				team.HeadCoach = row[9].(string)
				team.DLeagueAffiliation = row[10].(string)
			}
		}
	}
	return team, nil
}
