package domain

// A Politician is a leader whom we'd like to hold accountable to their Promises
type Politician struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Title      string    `json:"title"`
	CountryISO string    `json:"country" db:"country"`
	Promises   []Promise `json:"promises"`
}

// CreatePromise adds a new Promise for this Politician
func (p *Politician) CreatePromise(name string, details string, status PromiseStatus, category PromiseCategory, source Source) {
	p.Promises = append(p.Promises, Promise{
		PoliticianID: p.ID,
		Politician:   p,
		Name:         name,
		Details:      details,
		History:      []PromiseStatus{status},
		Category:     category,
		Sources:      []Source{source},
	})
}

// PoliticianRepository describes the interface for interacting with Politicians
type PoliticianRepository interface {
	All() ([]Politician, error)
	PoliticianOfIdentity(string) (*Politician, error)
	Save(*Politician) error

	SuggestionsOfPolitician(*Politician) ([]Suggestion, error)
	SaveSuggestion(*Suggestion) error
	AcceptSuggestion(*Suggestion) (*Promise, error)
	RejectSuggestion(*Suggestion) error
}
