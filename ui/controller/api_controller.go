package controller

import (
	"net/http"
	"strconv"

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
		l.errorReponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": politicians})
}

func (l *APIController) Promises(c *gin.Context) {
	promises, err := l.PoliticianSvc.ListPromisesOfPolitician(c.Param("politician"))
	if err != nil {
		l.errorReponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": promises})
}

func (l *APIController) Promise(c *gin.Context) {
	promiseID, err := strconv.Atoi(c.Param("promise"))
	if err != nil {
		l.errorReponse(err, c)
		return
	}

	promise, err := l.PoliticianSvc.SinglePromiseOfPolitician(c.Param("politician"), promiseID)
	if err != nil {
		l.errorReponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": promise})
}

func (l *APIController) errorReponse(err error, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "data": err.Error()})
}
