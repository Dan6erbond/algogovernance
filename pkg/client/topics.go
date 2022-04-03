package client

import "time"

// Topic represents a governance proposal topic.
type Topic struct {
	ID              int           `json:"id"`
	Title           string        `json:"title"`
	DescriptionHTML string        `json:"description_html"`
	TotalVoteCount  int           `json:"total_vote_count"`
	TopicOptions    []TopicOption `json:"topic_options"`
}

// TopicOption represents options that can be voted on within a topic.
type TopicOption struct {
	ID                 string  `json:"id"`
	Title              string  `json:"title"`
	Indicator          string  `json:"indicator"`
	IsFoundationChoice bool    `json:"is_foundation_choice"`
	VotePercentage     float64 `json:"vote_percentage,string"`
}

// TopicOptionVotes represents a list of votes for a topic option.
type TopicOptionVotes struct {
	Pagination
	Results []TopicOptionResults `json:"results"`
}

// GetNext returns the next page of results.
func (t *TopicOptionVotes) GetNext() (result TopicOptionVotes, err error) {
	if t.HasNext() {
		err = Get(t.Next, &result)
	} else {
		err = ErrNoNext
	}
	return result, err
}

// GetPrevious returns the previous page of results.
func (t *TopicOptionVotes) GetPrevious() (result TopicOptionVotes, err error) {
	if t.HasPrevious() {
		err = Get(t.Previous, &result)
	} else {
		err = ErrNoPrevious
	}
	return result, err
}

// TopicOptionResults contains details about a topic option vote, such as governor and allocated ALGO amount.
type TopicOptionResults struct {
	ID                    string                `json:"id"`
	GovernorVotingSession GovernorVotingSession `json:"governor_voting_session"`
	CreationDatetime      time.Time             `json:"creation_datetime"`
	AllocatedAlgoAmount   int                   `json:"allocated_algo_amount,string"`
}

// GovernorVotingSession includes details about the governor and transaction used to vote for a topic option.
type GovernorVotingSession struct {
	Governor    Governor    `json:"governor"`
	Transaction Transaction `json:"transaction"`
}
