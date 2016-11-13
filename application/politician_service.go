package application

import "github.com/politicalrev/accountability-api/domain"

type PoliticianService struct {
	PoliticianRepo domain.PoliticianRepository
}

func (s *PoliticianService) ListPoliticians() ([]domain.Politician, error) {
	return s.PoliticianRepo.All()
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
