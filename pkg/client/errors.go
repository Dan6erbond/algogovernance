package client

import "errors"

var (
	// ErrNoNext is returned when there is no next page of results.
	ErrNoNext     = errors.New("no next page")
	// ErrNoPrevious is returned when there is no previous page of results.
	ErrNoPrevious = errors.New("no previous page")
)

// ErrorResponse represents an error response from the Algorand API.
type ErrorResponse struct {
	Type            string              `json:"type"`   // Type is a human-readable error type.
	Detail          map[string][]string `json:"detail"` // Detail includes detailed error messages.
	FallbackMessage string              `json:"fallback_message"`
}
