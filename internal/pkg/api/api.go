package api

import (
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	r.GET("/players/:player_id", playerHandler)
	r.GET("/players", playersHandler)

	r.GET("/teams/:team_id", teamHandler)
	r.GET("/teams", teamsHandler)

	r.GET("/games", gamesHandler)

	r.Use(gin.Logger())

	r.Run(":8080")
}
