package routes

import (
	"github.com/t01gyl0p/scriptIq/handlers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {

	// Register route
	r.POST("/register", handlers.Register)

	//Login route
	r.POST("/login", handlers.LoginIn)

	// OAuth route
	r.POST("/auth/google/callback", handlers.OAuth)

	// Evaluate route
	r.POST("/evaluate", handlers.EvaluateCode)
}