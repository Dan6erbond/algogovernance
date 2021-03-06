package client

import "time"

// Transaction represents a transaction in the Algorand blockchain.
type Transaction struct {
	ID               string                      `json:"id"`
	TransactionID    string                      `json:"transaction_id"`
	RoundedAt        time.Time                   `json:"rounded_at"`
	ProcessedAt      time.Time                   `json:"processed_at"`
	GovernorActivity TransactionGovernorActivity `json:"governor_activity"`
}

// TransactionGovernorActivity includes information about the governor and the activity.
type TransactionGovernorActivity struct {
	ID               string      `json:"id"`
	Governor         Governor    `json:"governor"`
	ActivityType     string      `json:"activity_type"`
	CommittedAmount  interface{} `json:"committed_amount"`
	Message          string      `json:"message"`
	CreationDatetime time.Time   `json:"creation_datetime"`
}
