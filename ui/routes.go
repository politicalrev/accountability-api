package ui

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/politicalrev/accountability-api/infrastructure"
	"github.com/politicalrev/accountability-api/ui/middleware"
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
	apiV1 := router.Group("/v1")

	// Authentication requires a key/secret pair
	apiClientSvc := container.APIClientService()
	accounts, err := apiClientSvc.ListAccounts()
	if err != nil {
		log.Fatal(err)
	}
	apiV1.Use(gin.BasicAuth(accounts))

	// All api requests are logged
	apiV1.Use(middleware.LogRequest(apiClientSvc))

	// List available categories
	apiV1.GET("/categories", apiController.Categories)

	// List available statuses
	apiV1.GET("/statuses", apiController.Statuses)

	// List politicians tracked
	apiV1.GET("/politicians", apiController.Politicians)

	// List promises made by a politician
	apiV1.GET("/politicians/:politician/promises", apiController.Promises)

	// Get a single promise
	apiV1.GET("/politicians/:politician/promises/:promise", apiController.Promise)

	// See suggestions pending moderation
	apiV1.GET("/politicians/:politician/suggestions", apiController.Suggestions)

	// Submit promise data for moderation
	apiV1.POST("/politicians/:politician/suggestions", apiController.SubmitSuggestion)

	// Accept a suggestion
	apiV1.POST("/politicians/:politician/suggestions/:id/accept", apiController.AcceptSuggestion)

	// Reject a suggestion
	apiV1.POST("/politicians/:politician/suggestions/:id/reject", apiController.RejectSuggestion)
}
