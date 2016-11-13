package domain

// A Source is a citation for a Promise or update on the status
type Source struct {
	ID   string `json:"-"`
	Name string `json:"name"`
	Link string `json:"link"`
}
