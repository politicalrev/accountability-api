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
	politicianSvc *application.PoliticianService

	// Infrastructure
	db *sqlx.DB

	politicianRepo domain.PoliticianRepository

	// UI
	apiController *controller.APIController
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
