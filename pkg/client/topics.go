package client

import "time"

type Topic struct {
	ID              int           `json:"id"`
	Title           string        `json:"title"`
	DescriptionHTML string        `json:"description_html"`
	TotalVoteCount  int           `json:"total_vote_count"`
	TopicOptions    []TopicOption `json:"topic_options"`
}

type TopicOption struct {
	ID                 string  `json:"id"`
	Title              string  `json:"title"`
	Indicator          string  `json:"indicator"`
	IsFoundationChoice bool    `json:"is_foundation_choice"`
	VotePercentage     float64 `json:"vote_percentage,string"`
}

type TopicOptionVotes struct {
	Pagination
	Results []TopicOptionResults `json:"results"`
}

func (t *TopicOptionVotes) GetNext() (result TopicOptionVotes, err error) {
	if t.HasNext() {
		err = Get(t.Next, &result)
	} else {
		err = ErrNoNext
	}
	return result, err
}

func (t *TopicOptionVotes) GetPrevious() (result TopicOptionVotes, err error) {
	if t.HasPrevious() {
		err = Get(t.Previous, &result)
	} else {
		err = ErrNoPrevious
	}
	return result, err
}

type TopicOptionResults struct {
	ID                    string                `json:"id"`
	GovernorVotingSession GovernorVotingSession `json:"governor_voting_session"`
	CreationDatetime      time.Time             `json:"creation_datetime"`
	AllocatedAlgoAmount   int                   `json:"allocated_algo_amount,string"`
}

type GovernorVotingSession struct {
	Governor    Governor    `json:"governor"`
	Transaction Transaction `json:"transaction"`
}
