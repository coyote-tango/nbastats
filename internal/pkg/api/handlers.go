package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nbastats/internal/pkg/datamodel"
	"net/http"
	"strconv"
)

func getTeamHandler(c *gin.Context) {
	teamID, _ := strconv.Atoi(c.Param("team_id"))
	team, err := datamodel.GetTeam(teamID)
	fmt.Print(*team)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, team)
}

func getTeamsHandler(c *gin.Context) {
	teams, err := datamodel.GetTeams()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teams)
}

func getPlayersHandler(c *gin.Context) {
	players, err := datamodel.GetPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, players)
}

func getPlayerHandler(c *gin.Context) {
	playerID, _ := strconv.Atoi(c.Param("player_id"))
	player, err := datamodel.GetPlayer(playerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, player)
}
