package domain

import "github.com/gin-gonic/gin"

// APIClient is a consumer of this API
type APIClient struct {
	Key    string `db:"key"`
	Secret string `db:"secret"`
	Name   string `db:"name"`
}

// APIClientRepository allows access to the API Clients stored
type APIClientRepository interface {
	Accounts() (gin.Accounts, error)
	LogRequest(key, method, resource, ip string) error
}
