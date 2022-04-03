package client

import "time"

// AlgoCommitter allows structs to implement the GetCommittedAlgoAmount() method to return the committed algo amount for a governance period.
type AlgoCommitter interface {
	GetCommittedAlgoAmount() (committedAlgoAmount float64)
}

var _ AlgoCommitter = &Governor{}

// Governor represents a Algorand Governor associated with a wallet address and additional metadata such as committed algo amount for a governance period.
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

// GetCommittedAlgoAmount returns the committed algo amount for the current period to implement the AlgoCommitter interface.
func (g *Governor) GetCommittedAlgoAmount() (committedAlgoAmount float64) {
	return float64(g.CommittedAlgoAmount)
}

// Account represents a governor's account with ID and wallet address information.
type Account struct {
	ID      string `json:"id"`
	Address string `json:"address"`
}

var _ AlgoCommitter = &GovernorStatus{}

// GovernorStatus represents a Algorand Governor as returned by the /governors/%s/status endpoint.
type GovernorStatus struct {
	CommittedAlgoAmount float64 `json:"committed_algo_amount"`
	Period              struct {
		Title string `json:"title"`
	} `json:"period"`
}

// GetCommittedAlgoAmount returns the committed algo amount for the current period to implement the AlgoCommitter interface.
func (g *GovernorStatus) GetCommittedAlgoAmount() (committedAlgoAmount float64) {
	return g.CommittedAlgoAmount
}

// PeriodGovernorStatus represents an Algorand Governor's status for a governance period.
type PeriodGovernorStatus struct {
	Governor                  Governor        `json:"governor"`
	UncompletedVotingSessions []VotingSession `json:"uncompleted_voting_sessions"`
	Reward                    string          `json:"reward"`
}

// PeriodGovernor includes data about a governor's voting history in addition to the general governor data.
type PeriodGovernor struct {
	Governor
	VotingSessionHistory []PeriodGovernorVotingSessionHistory `json:"voting_session_history"`
}

// PeriodGovernorVotedOptions represents the voted option by a governor in a voting session.
type PeriodGovernorVotedOptions struct {
	ID                  string      `json:"id"`
	VotedOption         TopicOption `json:"voted_option"`
	AllocatedAlgoAmount string      `json:"allocated_algo_amount"`
}

// PeriodGovernorVotes contains a list of the votes cast by a governor in a voting session.
type PeriodGovernorVotes struct {
	ID               string                       `json:"id"`
	VotedOptions     []PeriodGovernorVotedOptions `json:"voted_options"`
	Topic            Topic                        `json:"topic"`
	CreationDatetime time.Time                    `json:"creation_datetime"`
}

// PeriodGovernorVotingSessionHistory contains a list of the voting sessions that a governor has voted in.
type PeriodGovernorVotingSessionHistory struct {
	VotingSession
	Votes []PeriodGovernorVotes `json:"votes"`
}

// GovernorActivites represents a list of an Algorand Governor's activities.
// This includes votes and stake commitments.
type GovernorActivities struct {
	Pagination
	Results []GovernorActivity `json:"results"`
}

// GetNext gets the next page of results if any are available.
// Otherwise throws a ErrNoNext error.
func (g *GovernorActivities) GetNext() (result GovernorActivities, err error) {
	if g.HasNext() {
		err = Get(g.Next, &result)
	} else {
		err = ErrNoNext
	}
	return result, err
}

// GetPrevious gets the previous page of results if any are available.
// Otherwise throws a ErrNoPrevious error.
func (g *GovernorActivities) GetPrevious() (result GovernorActivities, err error) {
	if g.HasPrevious() {
		err = Get(g.Previous, &result)
	} else {
		err = ErrNoPrevious
	}
	return result, err
}

// GovernorActivity represents an governance activities such as votes and stake commitments.
type GovernorActivity struct {
	ID               string      `json:"id"`
	ActivityType     string      `json:"activity_type"`
	CommittedAmount  string      `json:"committed_amount"`
	Message          string      `json:"message"`
	CreationDatetime time.Time   `json:"creation_datetime"`
	Transaction      Transaction `json:"transaction"`
}
