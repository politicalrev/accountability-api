package application

import (
	"github.com/gin-gonic/gin"
	"github.com/politicalrev/accountability-api/domain"
)

type APIClientService struct {
	APIClientRepo domain.APIClientRepository
}

// ListAccounts returns all registed API clients
func (s *APIClientService) ListAccounts() (gin.Accounts, error) {
	return s.APIClientRepo.Accounts()
}

// LogRequest logs an API request
func (s *APIClientService) LogRequest(key, method, resource, ip string) error {
	return s.APIClientRepo.LogRequest(key, method, resource, ip)
}
