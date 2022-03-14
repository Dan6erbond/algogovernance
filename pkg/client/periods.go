package client

import "time"

type Periods struct {
	Count    int                `json:"count"`
	Next     interface{}        `json:"next"`
	Previous interface{}        `json:"previous"`
	Results  []GovernancePeriod `json:"results"`
}

type GovernancePeriod struct {
	ID                      string           `json:"id"`
	Title                   string           `json:"title"`
	Slug                    string           `json:"slug"`
	SignUpAddress           string           `json:"sign_up_address"`
	RegistrationEndDatetime time.Time        `json:"registration_end_datetime"`
	StartDatetime           time.Time        `json:"start_datetime"`
	ActiveStateEndDatetime  time.Time        `json:"active_state_end_datetime"`
	EndDatetime             time.Time        `json:"end_datetime"`
	IsActive                bool             `json:"is_active"`
	TotalCommittedStake     float64          `json:"total_committed_stake"`
	AlgoAmountInRewardPool  int64            `json:"algo_amount_in_reward_pool,string"`
	GovernorCount           int              `json:"governor_count"`
	VotingSessions          []VotingSessions `json:"voting_sessions"`
}

type VotingSessions struct {
	ID                  int       `json:"id"`
	Title               string    `json:"title"`
	Slug                string    `json:"slug"`
	ShortDescription    string    `json:"short_description"`
	VotingStartDatetime time.Time `json:"voting_start_datetime"`
	VotingEndDatetime   time.Time `json:"voting_end_datetime"`
	TopicCount          int       `json:"topic_count"`
	CooldownEndDatetime time.Time `json:"cooldown_end_datetime"`
}
