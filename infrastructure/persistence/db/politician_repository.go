package db

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/politicalrev/accountability-api/domain"
)

type PoliticianRepository struct {
	DB *sqlx.DB
}

// All returns all politicians in the database
func (r *PoliticianRepository) All() ([]domain.Politician, error) {
	politicians := []domain.Politician{}
	if err := r.DB.Select(&politicians, "select * from politicians"); err != nil {
		return nil, err
	}

	// Shallow list (no promises loaded)
	return politicians, nil
}

// PoliticianOfName returns a specific politician by identity
func (r *PoliticianRepository) PoliticianOfIdentity(id string) (*domain.Politician, error) {
	politician := &domain.Politician{}
	if err := r.DB.Get(politician, "select * from politicians where id = $1", id); err != nil {
		return nil, err
	}

	// Populate with promise data
	if err := r.retrievePromises(politician); err != nil {
		return nil, err
	}

	return politician, nil
}

// Save persists the politician data to the database
func (r *PoliticianRepository) Save(*domain.Politician) error {
	return fmt.Errorf("Not implemented")
}

// SuggestionsOfPolitician returns all the suggestions submitted for a politician
func (r *PoliticianRepository) SuggestionsOfPolitician(p *domain.Politician) ([]domain.Suggestion, error) {
	suggestions := []domain.Suggestion{}
	if err := r.DB.Select(&suggestions, "select * from moderation_queue"); err != nil {
		return nil, err
	}

	return suggestions, nil
}

// SaveSuggestion persists a suggestion
func (r *PoliticianRepository) SaveSuggestion(s *domain.Suggestion) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	tx.Exec(`
        insert into moderation_queue (created_at, politician_id, promise, status, status_detail, category, source_name, source_link)
        values ($1, $2, $3, $4, $5, $6, $7, $8)
        `,
		time.Now(),
		s.PoliticianID,
		s.Promise,
		s.Status,
		s.StatusDetail,
		s.Category,
		s.SourceName,
		s.SourceLink,
	)

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *PoliticianRepository) retrievePromises(p *domain.Politician) error {
	rows, err := r.DB.Queryx("select * from promises where politician_id = $1", p.ID)
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		var prom domain.Promise
		if err := rows.StructScan(&prom); err != nil {
			return err
		}

		// Populate data
		prom.Politician = p
		if err := r.retrievePromiseHistory(&prom); err != nil {
			return err
		}

		if err := r.retrievePromiseSources(&prom); err != nil {
			return err
		}

		p.Promises = append(p.Promises, prom)
	}

	return nil
}

func (r *PoliticianRepository) retrievePromiseHistory(p *domain.Promise) error {
	rows, err := r.DB.Queryx("select * from promise_status where promise_id = $1 order by updated_on desc", p.ID)
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		var stat domain.PromiseStatus
		if err := rows.StructScan(&stat); err != nil {
			return err
		}

		stat.Promise = p
		if err := r.retrievePromiseStatusSources(&stat); err != nil {
			return err
		}

		p.History = append(p.History, stat)
	}

	return nil
}

func (r *PoliticianRepository) retrievePromiseSources(p *domain.Promise) error {
	rows, err := r.DB.Queryx("select s.* from promise_sources ps left join sources s on ps.source_id = s.id where ps.promise_id = $1", p.ID)
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		var source domain.Source
		if err := rows.StructScan(&source); err != nil {
			return err
		}

		p.Sources = append(p.Sources, source)
	}

	return nil
}

func (r *PoliticianRepository) retrievePromiseStatusSources(p *domain.PromiseStatus) error {
	rows, err := r.DB.Queryx("select s.* from promise_status_sources pss left join sources s on pss.source_id = s.id where pss.status_id = $1", p.ID)
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		var source domain.Source
		if err := rows.StructScan(&source); err != nil {
			return err
		}

		p.Sources = append(p.Sources, source)
	}

	return nil
}
