package infrastructure

import (
	"log"

	_ "github.com/lib/pq" // For sqlx binding

	"github.com/jmoiron/sqlx"
	"github.com/politicalrev/accountability-api/application"
	"github.com/politicalrev/accountability-api/domain"
	"github.com/politicalrev/accountability-api/infrastructure/persistence/db"
	"github.com/politicalrev/accountability-api/ui/controller"
	"github.com/spf13/viper"
)

// Container is a service locator
type Container struct {
	// Application
	apiClientSvc  *application.APIClientService
	politicianSvc *application.PoliticianService

	// Infrastructure
	db *sqlx.DB

	apiClientRepo  domain.APIClientRepository
	politicianRepo domain.PoliticianRepository

	// UI
	apiController *controller.APIController
}

// APIClientService returns an initialized application service for dealing with APIClients
func (c *Container) APIClientService() *application.APIClientService {
	if c.apiClientSvc == nil {
		c.apiClientSvc = &application.APIClientService{
			APIClientRepo: c.apiClientRepository(),
		}
	}

	return c.apiClientSvc
}

// PoliticianService returns an initialized application service for dealing with Politicians
func (c *Container) PoliticianService() *application.PoliticianService {
	if c.politicianSvc == nil {
		c.politicianSvc = &application.PoliticianService{
			PoliticianRepo: c.politicianRepository(),
		}
	}

	return c.politicianSvc
}

func (c *Container) database() *sqlx.DB {
	if c.db == nil {
		db, err := sqlx.Connect("postgres", viper.GetString("database_url"))
		if err != nil {
			log.Fatalln(err)
		}

		c.db = db
	}

	return c.db
}

func (c *Container) apiClientRepository() domain.APIClientRepository {
	if c.apiClientRepo == nil {
		c.apiClientRepo = &db.APIClientRepository{
			DB: c.database(),
		}
	}

	return c.apiClientRepo
}

func (c *Container) politicianRepository() domain.PoliticianRepository {
	if c.politicianRepo == nil {
		c.politicianRepo = &db.PoliticianRepository{
			DB: c.database(),
		}
	}

	return c.politicianRepo
}

// APIController returns an initialized controller for /api routes
func (c *Container) APIController() *controller.APIController {
	if c.apiController == nil {
		c.apiController = &controller.APIController{
			PoliticianSvc: c.PoliticianService(),
		}
	}

	return c.apiController
}
