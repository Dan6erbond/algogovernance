package client

import "time"

type AlgoCommitter interface {
	GetCommittedAlgoAmount() (committedAlgoAmount float64)
}

var _ AlgoCommitter = &Governor{}

type Governor struct {
	ID                      string    `json:"id"`
	Account                 Account   `json:"account"`
	BeneficiaryAccount      Account   `json:"beneficiary_account"`
	CommittedAlgoAmount     int       `json:"committed_algo_amount,string"`
	IsEligible              bool      `json:"is_eligible"`
	NotEligibleReason       string    `json:"not_eligible_reason"`
	RegistrationDatetime    time.Time `json:"registration_datetime"`
	VotedVotingSessionCount int       `json:"voted_voting_session_count"`
}

func (g *Governor) GetCommittedAlgoAmount() (committedAlgoAmount float64) {
	return float64(g.CommittedAlgoAmount)
}

type Account struct {
	ID      string `json:"id"`
	Address string `json:"address"`
}

var _ AlgoCommitter = &GovernorStatus{}

type GovernorStatus struct {
	CommittedAlgoAmount float64 `json:"committed_algo_amount"`
	Period              struct {
		Title string `json:"title"`
	} `json:"period"`
}

func (g *GovernorStatus) GetCommittedAlgoAmount() (committedAlgoAmount float64) {
	return g.CommittedAlgoAmount
}

type PeriodGovernorStatus struct {
	Governor                  Governor        `json:"governor"`
	UncompletedVotingSessions []VotingSession `json:"uncompleted_voting_sessions"`
	Reward                    string          `json:"reward"`
}

type PeriodGovernor struct {
	Governor
	VotingSessionHistory []PeriodGovernorVotingSessionHistory `json:"voting_session_history"`
}

type PeriodGovernorVotedOptions struct {
	ID                  string      `json:"id"`
	VotedOption         TopicOption `json:"voted_option"`
	AllocatedAlgoAmount string      `json:"allocated_algo_amount"`
}

type PeriodGovernorVotes struct {
	ID               string                       `json:"id"`
	VotedOptions     []PeriodGovernorVotedOptions `json:"voted_options"`
	Topic            Topic                        `json:"topic"`
	CreationDatetime time.Time                    `json:"creation_datetime"`
}

type PeriodGovernorVotingSessionHistory struct {
	VotingSession
	Votes []PeriodGovernorVotes `json:"votes"`
}

type GovernorActivities struct {
	Pagination
	Results []GovernorActivity `json:"results"`
}

func (g *GovernorActivities) GetNext() (result GovernorActivities, err error) {
	if g.HasNext() {
		err = Get(g.Next, &result)
	} else {
		err = ErrNoNext
	}
	return result, err
}

func (g *GovernorActivities) GetPrevious() (result GovernorActivities, err error) {
	if g.HasPrevious() {
		err = Get(g.Previous, &result)
	} else {
		err = ErrNoPrevious
	}
	return result, err
}

type GovernorActivity struct {
	ID               string      `json:"id"`
	ActivityType     string      `json:"activity_type"`
	CommittedAmount  string      `json:"committed_amount"`
	Message          string      `json:"message"`
	CreationDatetime time.Time   `json:"creation_datetime"`
	Transaction      Transaction `json:"transaction"`
}
