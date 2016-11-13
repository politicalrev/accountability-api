package domain

// PromiseCategory is the general topic that the Promise falls into
type PromiseCategory string

const (
	// Climate - Any Promise relating to climate and the environment
	Climate PromiseCategory = "climate"

	// Culture - Any Promise relating to lifestyle
	Culture PromiseCategory = "culture"

	// Economy - Any Promise relating to the economy and taxes
	Economy PromiseCategory = "economy"

	// Government - Any Promise relating to government reform
	Government PromiseCategory = "government"

	// Healthcare - Any Promise relating to healthcare concerns
	Healthcare PromiseCategory = "healthcare"

	// Immigration - Any Promise relating to immigration and visas
	Immigration PromiseCategory = "immigration"

	// Security - Any Promise relating to war and defense
	Security PromiseCategory = "security"
)

// A Promise is a declaration that a Politician makes about something they will do
type Promise struct {
	ID           int             `json:"id"`
	PoliticianID string          `json:"-" db:"politician_id"`
	Politician   *Politician     `json:"-"`
	Name         string          `json:"name"`
	Details      string          `json:"details"`
	History      []PromiseStatus `json:"history"`
	Category     PromiseCategory `json:"category" db:"category"`
	Sources      []Source        `json:"sources"`
}
