package api

import (
	"github.com/gin-gonic/gin"
	// Import other required packages, like your internal NBA Stats package
)

func StartServer() {
	r := gin.Default()

	// Define your API routes and handlers
	r.GET("/teams", getTeamsHandler)
	r.GET("/teams/:team_id", getTeamHandler)

	// Start the HTTP server
	r.Run(":8080")
}
