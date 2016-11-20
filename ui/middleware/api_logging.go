package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/politicalrev/accountability-api/application"
)

// LogRequest logs all API requests
func LogRequest(s *application.APIClientService) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _, _ := c.Request.BasicAuth()
		if err := s.LogRequest(user, c.Request.Method, c.Request.URL.String(), c.ClientIP()); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		c.Next()
	}
}
