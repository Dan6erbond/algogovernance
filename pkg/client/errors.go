package client

import "errors"

var (
	ErrNoNext     = errors.New("no next page")
	ErrNoPrevious = errors.New("no previous page")
)

// ErrorResponse represents an error response from the Algorand API.
type ErrorResponse struct {
	Type            string              `json:"type"`   // Type is a human-readable error type.
	Detail          map[string][]string `json:"detail"` // Detail includes detailed error messages.
	FallbackMessage string              `json:"fallback_message"`
}
