package client

// Pagination represents pagination information for a list of resources.
// It is intended to be embedded in a struct that represents a list of resources.
type Pagination struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

// HasNext returns true if there is a next page of results.
func (p *Pagination) HasNext() bool {
	return p.Next != ""
}

// HasPrevious returns true if there is a previous page of results.
func (p *Pagination) HasPrevious() bool {
	return p.Previous != ""
}
