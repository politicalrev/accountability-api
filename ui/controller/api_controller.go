package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/politicalrev/accountability-api/application"
)

type APIController struct {
	PoliticianSvc *application.PoliticianService
}

func (l *APIController) Version(c *gin.Context) {
	c.String(http.StatusOK, "1")
}

func (l *APIController) Politicians(c *gin.Context) {
	politicians, err := l.PoliticianSvc.ListPoliticians()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": politicians})
}

func (l *APIController) Promises(c *gin.Context) {
	promises, err := l.PoliticianSvc.ListPromisesOfPolitician(c.Param("politician"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": promises})
}
