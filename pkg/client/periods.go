package client

import "time"

// Period represents a governance period and may not contain all fields depending on the endpoint used.
// See more at https://governance.algorand.foundation/api/documentation/.
type Period struct {
	ID                      string          `json:"id"`
	Title                   string          `json:"title"`
	Slug                    string          `json:"slug"`
	SignUpAddress           string          `json:"sign_up_address"`
	RegistrationEndDatetime time.Time       `json:"registration_end_datetime"`
	StartDatetime           time.Time       `json:"start_datetime"`
	ActiveStateEndDatetime  time.Time       `json:"active_state_end_datetime"`
	EndDatetime             time.Time       `json:"end_datetime"`
	IsActive                bool            `json:"is_active"`
	TotalCommittedStake     float64         `json:"total_committed_stake"`
	AlgoAmountInRewardPool  int             `json:"algo_amount_in_reward_pool,string"`
	GovernorCount           int             `json:"governor_count"`
	VotingSessions          []VotingSession `json:"voting_sessions"`
}

// PeriodStatistics general statistics of the governance platform.
type PeriodStatistics struct {
	UniqueGovernorsCount    int      `json:"unique_governors_count"`
	TotalRewardsDistributed float64  `json:"total_rewards_distributed"`
	PastPeriods             []Period `json:"past_periods"`
}

// Periods contains a list of periods and can be paginated using cursor pagination.
type Periods struct {
	Pagination
	Results []Period `json:"results"`
}

// GetNext returns the next page of results.
func (p *Periods) GetNext() (result Periods, err error) {
	if p.HasNext() {
		err = Get(p.Next, &result)
	} else {
		err = ErrNoNext
	}
	return result, err
}

// GetPrevious returns the previous page of results.
func (p *Periods) GetPrevious() (result Periods, err error) {
	if p.HasPrevious() {
		err = Get(p.Previous, &result)
	} else {
		err = ErrNoPrevious
	}
	return result, err
}

// PeriodGovernors contains a list of governors and pagination information.
type PeriodGovernors struct {
	Pagination
	Results []Governor `json:"results"`
}

// GetNext returns the next page of results.
func (p *PeriodGovernors) GetNext() (result PeriodGovernors, err error) {
	if p.HasNext() {
		err = Get(p.Next, &result)
	} else {
		err = ErrNoNext
	}
	return result, err
}

// GetPrevious returns the previous page of results.
func (p *PeriodGovernors) GetPrevious() (result PeriodGovernors, err error) {
	if p.HasPrevious() {
		err = Get(p.Previous, &result)
	} else {
		err = ErrNoPrevious
	}
	return result, err
}
