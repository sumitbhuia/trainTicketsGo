package routes

import (
	"trainTicketsGo/controllers"

	"github.com/gin-gonic/gin"
)

// UserRoutes function
// All routes for /api/v1/trains
//Basically "r" is similar to the router object in Express.js
// This function takes a router object (r) and adds routes to it.
// The routes are defined in the controllers package.

func UserRoutes(r *gin.Engine) {
	r.POST("/api/v1/auth/register", controllers.Register) // Sign up new user
	r.POST("/api/v1/auth/login", controllers.Login)       // Login user
}
