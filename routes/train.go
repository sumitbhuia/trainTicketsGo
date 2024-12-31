package routes

import (
	"trainTicketsGo/controllers"
	"trainTicketsGo/middlewares"

	"github.com/gin-gonic/gin"
)

// TrainRoutes function
// All routes for /api/v1/trains
//Basically "r" is similar to the router object in Express.js
// This function takes a router object (r) and adds routes to it.
// The routes are defined in the controllers package.

func TrainRoutes(r *gin.Engine) {

	// Protected routes
	secured := r.Group("/")
	secured.Use(middlewares.AuthMiddleware())
	{
		secured.GET("/api/v1/trains", controllers.GetTrains) // Get all trains
		secured.POST("/api/v1/trains", controllers.AddTrain) // Add a train
	}
}
