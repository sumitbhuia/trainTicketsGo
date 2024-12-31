package main

import (
	"trainTicketsGo/database"
	"trainTicketsGo/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// Connect to the database
	database.ConnectDB()

	// Setup routes
	// Create a new gin router and pass it to the router function
	r := gin.Default()
	routes.TrainRoutes(r)
	routes.TicketRoutes(r)
	routes.UserRoutes(r)

	// Start the server
	r.Run(":3000")
}
