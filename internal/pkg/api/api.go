package api

import (
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	r.GET("/teams", getTeamsHandler)
	r.GET("/teams/:team_id", getTeamHandler)

	r.GET("/players/:player_id", getPlayerHandler)
	r.GET("/players", getPlayersHandler)

	r.Use(gin.Logger())

	r.Run(":8080")
}
