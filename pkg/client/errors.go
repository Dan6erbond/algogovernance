package client

import "errors"

var (
	ErrNoNext     = errors.New("no next page")
	ErrNoPrevious = errors.New("no previous page")
)

type ErrorResponse struct {
	Type            string              `json:"type"`
	Detail          map[string][]string `json:"detail"`
	FallbackMessage string              `json:"fallback_message"`
}
