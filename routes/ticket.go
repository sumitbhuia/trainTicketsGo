package routes

import (
	"trainTicketsGo/controllers"
	"trainTicketsGo/middlewares"

	"github.com/gin-gonic/gin"
)

// TicketRoutes function
// All routes for /api/v1/trains
//Basically "r" is similar to the router object in Express.js
// This function takes a router object (r) and adds routes to it.
// The routes are defined in the controllers package.

func TicketRoutes(r *gin.Engine) {

	// Protected routes
	secured := r.Group("/")
	secured.Use(middlewares.AuthMiddleware())
	{
		secured.POST("/api/v1/tickets", controllers.BookTicket)          // Book a ticket
		secured.POST("/api/v1/tickets/cancel", controllers.CancelTicket) // Cancel a ticket
	}
}
