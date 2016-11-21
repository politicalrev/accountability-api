package controller

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/politicalrev/accountability-api/application"
)

type APIController struct {
	PoliticianSvc *application.PoliticianService
}

func (l *APIController) Politicians(c *gin.Context) {
	politicians, err := l.PoliticianSvc.ListPoliticians()
	if err != nil {
		l.serverErrorReponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": politicians})
}

func (l *APIController) Categories(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": l.PoliticianSvc.ListCategories()})
}

func (l *APIController) Statuses(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": l.PoliticianSvc.ListStatuses()})
}

func (l *APIController) Promises(c *gin.Context) {
	promises, err := l.PoliticianSvc.ListPromisesOfPolitician(c.Param("politician"))
	if err != nil {
		l.serverErrorReponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": promises})
}

func (l *APIController) Promise(c *gin.Context) {
	promiseID, err := strconv.Atoi(c.Param("promise"))
	if err != nil {
		l.serverErrorReponse(err, c)
		return
	}

	promise, err := l.PoliticianSvc.SinglePromiseOfPolitician(c.Param("politician"), promiseID)
	if err != nil {
		l.serverErrorReponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": promise})
}

func (l *APIController) Suggestions(c *gin.Context) {
	suggestions, err := l.PoliticianSvc.ListSuggestions(c.Param("politician"))
	if err != nil {
		l.serverErrorReponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": suggestions})
}

func (l *APIController) SubmitSuggestion(c *gin.Context) {
	// Pull data from request body
	type promiseSubmission struct {
		Promise      string `json:"promise" valid:"required"`
		Status       string `json:"status" valid:"required,status"`
		StatusDetail string `json:"status_detail" valid:"required"`
		Category     string `json:"category" valid:"required,category"`
		SourceName   string `json:"source_name" valid:"required"`
		SourceLink   string `json:"source_link" valid:"required,url"`
	}

	var sub promiseSubmission
	if err := c.BindJSON(&sub); err != nil {
		l.requestErrorReponse(err, c)
		return
	}

	// Validate
	govalidator.TagMap["status"] = govalidator.Validator(func(s string) bool {
		for _, status := range l.PoliticianSvc.ListStatuses() {
			if s == string(status) {
				return true
			}
		}
		return false
	})

	govalidator.TagMap["category"] = govalidator.Validator(func(c string) bool {
		for _, category := range l.PoliticianSvc.ListCategories() {
			if c == string(category) {
				return true
			}
		}
		return false
	})

	_, err := govalidator.ValidateStruct(sub)
	if err != nil {
		l.requestErrorReponse(err, c)
		return
	}

	// Submit for moderation
	err = l.PoliticianSvc.SubmitSuggestionForModeration(
		c.Param("politician"),
		sub.Promise,
		sub.Status,
		sub.StatusDetail,
		sub.Category,
		sub.SourceName,
		sub.SourceLink,
	)
	if err != nil {
		l.serverErrorReponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (l *APIController) AcceptSuggestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		l.requestErrorReponse(err, c)
		return
	}

	promise, err := l.PoliticianSvc.AcceptSuggestion(c.Param("politician"), id, c.Query("user"))
	if err != nil {
		l.serverErrorReponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": promise})
}

func (l *APIController) RejectSuggestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		l.requestErrorReponse(err, c)
		return
	}

	err = l.PoliticianSvc.RejectSuggestion(c.Param("politician"), id, c.Query("user"))
	if err != nil {
		l.serverErrorReponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (l *APIController) requestErrorReponse(err error, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "error", "data": err.Error()})
}

func (l *APIController) serverErrorReponse(err error, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "data": err.Error()})
}
