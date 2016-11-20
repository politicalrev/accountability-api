package db

import (
	"fmt"

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
	if err := r.DB.Select(&suggestions, "select * from moderation_queue where accepted_at is null and deleted_at is null"); err != nil {
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
        values (NOW(), $1, $2, $3, $4, $5, $6, $7)
        `,
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

// AcceptSuggestion converts a suggestion to a promise
func (r *PoliticianRepository) AcceptSuggestion(s *domain.Suggestion) (*domain.Promise, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return nil, err
	}

	// Mark the suggestion as accepted
	tx.Exec(`
        update moderation_queue
        set accepted_at = NOW(),
            accepted_by = $1
        where
            id = $2
        `,
		s.AcceptedBy,
		s.ID,
	)

	// Insert the promise
	tx.Exec(`
        insert into promises (politician_id, name, category)
        values ($1, $2, $3)
        `,
		s.PoliticianID,
		s.Promise,
		s.Category,
	)

	// And status
	tx.Exec(`
        insert into promise_status (promise_id, name, updated_on, detail)
        select currval('promises_id_seq'), $1, NOW(), $2
        `,
		s.Status,
		s.StatusDetail,
	)

	// And source
	tx.Exec(`
        insert into sources (name, link)
        values ($1, $2)
        `,
		s.SourceName,
		s.SourceLink,
	)

	tx.Exec(`
        insert into promise_sources (promise_id, source_id)
        select currval('promises_id_seq'), currval('sources_id_seq')
        `,
	)

	tx.Exec(`
        insert into promise_status_sources (status_id, source_id)
        select currval('promise_status_id_seq'), currval('sources_id_seq')
        `,
	)

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	var promise domain.Promise
	if err := r.DB.Get(&promise, "select * from promises order by id desc limit 1"); err != nil {
		return nil, err
	}

	return &promise, nil
}

// RejectSuggestion removes a suggestion
func (r *PoliticianRepository) RejectSuggestion(s *domain.Suggestion) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	tx.Exec(`
        update moderation_queue
        set deleted_at = NOW(),
            deleted_by = $1
        where
            id = $2
        `,
		s.DeletedBy,
		s.ID,
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
