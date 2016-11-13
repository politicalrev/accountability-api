package ui

import (
	"net/http"

	"github.com/politicalrev/accountability-api/infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	container := infrastructure.Container{}

	// Health check
	router.HEAD("/ping", func(c *gin.Context) {})
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// API routes
	apiController := container.APIController()
	apiV1 := router.Group("/api")
	{
		apiV1.GET("/version", apiController.Version)

		// List politicians tracked
		apiV1.GET("/politicians", apiController.Politicians)

		// List promises made by a politician
		apiV1.GET("/politicians/:politician/promises", apiController.Promises)

		// Get a single promise
		apiV1.GET("/politicians/:politician/promises/:promise", apiController.Promise)
	}
}
