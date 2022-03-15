package client

type Pagination struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

func (p *Pagination) HasNext() bool {
	return p.Next != ""
}

func (p *Pagination) HasPrevious() bool {
	return p.Previous != ""
}
