package client

type ErrorResponse struct {
	Type            string              `json:"type"`
	Detail          map[string][]string `json:"detail"`
	FallbackMessage string              `json:"fallback_message"`
}
