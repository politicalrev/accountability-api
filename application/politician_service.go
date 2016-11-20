package application

import (
	"database/sql"
	"fmt"

	"github.com/politicalrev/accountability-api/domain"
)

type PoliticianService struct {
	PoliticianRepo domain.PoliticianRepository
}

func (s *PoliticianService) ListPoliticians() ([]domain.Politician, error) {
	return s.PoliticianRepo.All()
}

func (s *PoliticianService) ListCategories() []domain.PromiseCategory {
	return domain.ValidCategories()
}

func (s *PoliticianService) ListStatuses() []domain.PromiseStatusName {
	return domain.ValidStatuses()
}

func (s *PoliticianService) PoliticianExists(id string) (bool, error) {
	// It's actually cheaper to pull a shallow list of all politicians and iterate
	politicians, err := s.PoliticianRepo.All()
	if err != nil {
		return false, err
	}

	for _, politician := range politicians {
		if politician.ID == id {
			return true, nil
		}
	}

	return false, nil
}

func (s *PoliticianService) ListPromisesOfPolitician(id string) ([]domain.Promise, error) {
	politician, err := s.PoliticianRepo.PoliticianOfIdentity(id)
	if err != nil {
		return nil, err
	}

	return politician.Promises, nil
}

func (s *PoliticianService) SinglePromiseOfPolitician(politicianID string, id int) (*domain.Promise, error) {
	politician, err := s.PoliticianRepo.PoliticianOfIdentity(politicianID)
	if err != nil {
		return nil, err
	}

	for _, promise := range politician.Promises {
		if promise.ID == id {
			return &promise, nil
		}
	}

	return nil, nil
}

func (s *PoliticianService) ListSuggestions(id string) ([]domain.Suggestion, error) {
	politician, err := s.PoliticianRepo.PoliticianOfIdentity(id)
	if err != nil {
		return nil, err
	}

	return s.PoliticianRepo.SuggestionsOfPolitician(politician)
}

func (s *PoliticianService) SubmitSuggestionForModeration(politicianID, promise, status, statusDetail, category, sourceName, sourceLink string) error {
	return s.PoliticianRepo.SaveSuggestion(&domain.Suggestion{
		PoliticianID: politicianID,
		Promise:      promise,
		Status:       status,
		StatusDetail: statusDetail,
		Category:     category,
		SourceName:   sourceName,
		SourceLink:   sourceLink,
	})
}

func (s *PoliticianService) AcceptSuggestion(politicianID string, suggestionID int, userID string) (*domain.Promise, error) {
	suggestions, err := s.ListSuggestions(politicianID)
	if err != nil {
		return nil, err
	}

	for _, sugg := range suggestions {
		if sugg.ID == suggestionID {
			sugg.AcceptedBy = toNullString(userID)
			return s.PoliticianRepo.AcceptSuggestion(&sugg)
		}
	}

	return nil, fmt.Errorf("No suggestion of ID %d", suggestionID)
}

func (s *PoliticianService) RejectSuggestion(politicianID string, suggestionID int, userID string) error {
	suggestions, err := s.ListSuggestions(politicianID)
	if err != nil {
		return err
	}

	for _, sugg := range suggestions {
		if sugg.ID == suggestionID {
			sugg.DeletedBy = toNullString(userID)
			return s.PoliticianRepo.RejectSuggestion(&sugg)
		}
	}

	return fmt.Errorf("No suggestion of ID %d", suggestionID)
}

func toNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}
