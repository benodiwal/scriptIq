package main

import (
	 "github.com/t01gyl0p/scriptIq/middlewares"
	 "github.com/t01gyl0p/scriptIq/models"
	 "github.com/t01gyl0p/scriptIq/routes"
	 "github.com/gin-gonic/gin"
	 "github.com/gin-contrib/cors"
)

func main() {

	// Connecting database
	models.ConnectDatabase()

	// Initialize router
	r := gin.Default()

	// use CORS middleware
	r.Use(cors.Default())

	// Logger middleware
	r.Use(middlewares.LoggerMiddleware())

	// Initialize routes
	routes.InitializeRoutes(r)

	// Starting server on port 3000
	r.Run(":3001")
}