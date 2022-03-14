package client

import "time"

type Governors struct {
	Governor                  Governor      `json:"governor"`
	UncompletedVotingSessions []interface{} `json:"uncompleted_voting_sessions"`
	Reward                    interface{}   `json:"reward"`
}

type Account struct {
	ID      string `json:"id"`
	Address string `json:"address"`
}

type Governor struct {
	ID                      string      `json:"id"`
	Account                 Account     `json:"account"`
	BeneficiaryAccount      interface{} `json:"beneficiary_account"`
	CommittedAlgoAmount     int64       `json:"committed_algo_amount,string"`
	IsEligible              bool        `json:"is_eligible"`
	NotEligibleReason       string      `json:"not_eligible_reason"`
	RegistrationDatetime    time.Time   `json:"registration_datetime"`
	VotedVotingSessionCount int         `json:"voted_voting_session_count"`
}
