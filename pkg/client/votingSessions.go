package client

import "time"

// VotingSession represents a voting session in the Algorand blockchain.
type VotingSession struct {
	ID                  int       `json:"id"`
	Title               string    `json:"title"`
	Slug                string    `json:"slug"`
	ShortDescription    string    `json:"short_description"`
	VotingStartDatetime time.Time `json:"voting_start_datetime"`
	VotingEndDatetime   time.Time `json:"voting_end_datetime"`
	TopicCount          int       `json:"topic_count"`
	CooldownEndDatetime time.Time `json:"cooldown_end_datetime"`
}

// VotingSessionDetail contains additional details about a voting session as returned by the voting-sessions/%s endpoint.
type VotingSessionDetail struct {
	VotingSession
	DescriptionHTML                     string    `json:"description_html"`
	PublishedDatetime                   time.Time `json:"published_datetime"`
	Topics                              []Topic   `json:"topics"`
	TotalVotedGovernorsCount            int       `json:"total_voted_governors_count"`
	TotalCommittedStakeOfVotedGovernors float64   `json:"total_committed_stake_of_voted_governors"`
	TotalGovernorCount                  int       `json:"total_governor_count"`
	TotalCommittedStake                 float64   `json:"total_committed_stake"`
	HasFoundationProposal               bool      `json:"has_foundation_proposal"`
}
