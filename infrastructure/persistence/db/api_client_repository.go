package db

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/politicalrev/accountability-api/domain"
)

// APIClientRepository concrete implementation for the database
type APIClientRepository struct {
	DB *sqlx.DB
}

// Accounts returns all API clients registered
func (r *APIClientRepository) Accounts() (gin.Accounts, error) {
	clients := []domain.APIClient{}
	if err := r.DB.Select(&clients, "select * from api_clients"); err != nil {
		return nil, err
	}

	accounts := gin.Accounts{}
	for _, client := range clients {
		accounts[client.Key] = client.Secret
	}

	return accounts, nil
}

// LogRequest logs an API request in the database
func (r *APIClientRepository) LogRequest(key, method, resource, ip string) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	tx.Exec(`
        insert into api_requests (key, method, resource, requested_at, requested_by)
        values ($1, $2, $3, NOW(), $4)
        `,
		key,
		method,
		resource,
		ip,
	)

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
