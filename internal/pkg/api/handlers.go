package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nbastats/internal/pkg/datamodel"
	"net/http"
	"strconv"
	"time"
)

func playerHandler(c *gin.Context) {
	playerID, _ := strconv.Atoi(c.Param("player_id"))
	player, err := datamodel.GetPlayer(playerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, player)
}

func playersHandler(c *gin.Context) {
	players, err := datamodel.GetPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, players)
}

func teamHandler(c *gin.Context) {
	teamID, _ := strconv.Atoi(c.Param("team_id"))
	team, err := datamodel.GetTeam(teamID)
	fmt.Print(*team)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, team)
}

func teamsHandler(c *gin.Context) {
	teams, err := datamodel.GetTeams()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teams)
}

func gamesHandler(c *gin.Context) {
	// Grab gameDate from query string, if not provided, use time.Now() as default
	gameDate := c.DefaultQuery("gameDate", time.Now().Format("2006-01-02"))
	parsedGameDate, err := time.Parse("2006-01-02", gameDate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format, please use YYYY-MM-DD"})
		return
	}
	//TODO if date < today show score of previous games
	games, err := datamodel.GetGames(parsedGameDate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, games)
}
