package client

import "testing"

func TestGetPeriods(t *testing.T) {
	periods, err := GetPeriods("0", "0")
	if err != nil {
		t.Error(err)
	}
	if periods.Count == 0 {
		t.Error("Expected at least one period")
	}
	for _, period := range periods.Results {
		if period.ID == "" {
			t.Error("Expected period ID")
		}
		if period.Slug == "" {
			t.Error("Expected period slug")
		}
		if period.StartDatetime.IsZero() {
			t.Error("Expected period start datetime")
		}
		if period.EndDatetime.IsZero() {
			t.Error("Expected period end datetime")
		}
	}
}

func TestGetActivePeriod(t *testing.T) {
	period, err := GetActivePeriod()
	if err != nil {
		t.Error(err)
	}
	if period.ID == "" {
		t.Error("Expected active period ID")
	}
	if period.Slug == "" {
		t.Error("Expected active period slug")
	}
}

func TestGetPeriodStatistics(t *testing.T) {
	periodStatistics, err := GetPeriodStatistics()
	if err != nil {
		t.Error(err)
	}
	if periodStatistics.UniqueGovernorsCount == 0 {
		t.Error("Expected unique governors count")
	}
	if periodStatistics.TotalRewardsDistributed == 0 {
		t.Error("Expected total rewards distributed")
	}
	if len(periodStatistics.PastPeriods) == 0 {
		t.Error("Expected past periods")
	}
}

func TestGetPeriod(t *testing.T) {
	periods, err := GetPeriods("0", "0")
	if err != nil {
		t.Error(err)
	}
	if periods.Count == 0 {
		t.Error("Expected at least one period")
	}
	for _, p := range periods.Results {
		period, err := GetPeriod(p.Slug)
		if err != nil {
			t.Error(err)
		}
		if period.ID != p.ID {
			t.Errorf("Expected period ID to be %s, got %s", p.ID, period.ID)
		}
		if period.Slug != p.Slug {
			t.Errorf("Expected period slug to be %s, got %s", p.Slug, period.Slug)
		}
		if period.StartDatetime != p.StartDatetime {
			t.Errorf("Expected period start datetime to be %s, got %s", p.StartDatetime, period.StartDatetime)
		}
		if period.EndDatetime != p.EndDatetime {
			t.Errorf("Expected period end datetime to be %s, got %s", p.EndDatetime, period.EndDatetime)
		}
		if period.TotalCommittedStake != p.TotalCommittedStake {
			t.Errorf("Expected period total committed stake to be %.2f, got %.2f", p.TotalCommittedStake, period.TotalCommittedStake)
		}
		//lint:ignore SA4004 Only getting the first result
		break
	}
}

func TestGetPeriodGovernors(t *testing.T) {
	slug := "governance-period-2"
	periodGovernors, err := GetPeriodGovernors(slug, "", "", "", "", "", "", "")
	if err != nil {
		t.Error(err)
	}
	if periodGovernors.Count == 0 || len(periodGovernors.Results) == 0 {
		t.Error("Expected at least one period governor")
	}
	for _, periodGovernor := range periodGovernors.Results {
		if periodGovernor.ID == "" {
			t.Error("Expected period governor ID")
		}
		if periodGovernor.Account.ID == "" {
			t.Error("Expected period governor account ID")
		}
		if periodGovernor.Account.Address == "" {
			t.Error("Expected period governor account address")
		}
	}
	if periodGovernors.HasNext() {
		nextPage, err := periodGovernors.GetNext()
		if err != nil {
			t.Error(err)
		}
		if len(nextPage.Results) == 0 {
			t.Error("Expected more period governors")
		}
	}
}

func TestGetPeriodGovernor(t *testing.T) {
	slug := "governance-period-2"
	address := "3RYOY2LTPC6GLT3ZYE4LUFGGAEMY7GRENZQO7RFNGK2LGCV77QNASK6C6Y"
	periodGovernor, err := GetPeriodGovernor(slug, address)
	if err != nil {
		t.Error(err)
	}
	if periodGovernor.ID == "" {
		t.Error("Expected period governor ID")
	}
	if periodGovernor.CommittedAlgoAmount != 170000000 {
		t.Errorf("Expected period governor committed ALGO amount to be %d, got %d", 170000000, periodGovernor.CommittedAlgoAmount)
	}
	if periodGovernor.VotedVotingSessionCount != 1 {
		t.Errorf("Expected period governor voted voting session count to be %d, got %d", 1, periodGovernor.VotedVotingSessionCount)
	}
}

func TestGetGovernorActivities(t *testing.T) {
	slug := "governance-period-2"
	address := "3RYOY2LTPC6GLT3ZYE4LUFGGAEMY7GRENZQO7RFNGK2LGCV77QNASK6C6Y"
	governorActivities, err := GetGovernorActivities(slug, address, "", "")
	if err != nil {
		t.Error(err)
	}
	if governorActivities.Count != 2 {
		t.Errorf("Expected 2 governor activities, got %d", governorActivities.Count)
	}
}

func TestGetGovernorStatus(t *testing.T) {
	address := "3RYOY2LTPC6GLT3ZYE4LUFGGAEMY7GRENZQO7RFNGK2LGCV77QNASK6C6Y"
	_, err := GetGovernorStatus(address)
	if err != nil {
		t.Error(err)
	}
	// Cannot test additional governor status fields as they may be empty depending on the governance period
}

func TestGetTopicOptionVotes(t *testing.T) {
	votingSessionSlug := "voting-session-q1-2022"
	votingSession, err := GetVotingSession(votingSessionSlug)
	if err != nil {
		t.Error(err)
	}
	for _, topic := range votingSession.Topics {
		for _, option := range topic.TopicOptions {
			votes, err := GetTopicOptionVotes(option.ID, "", "")
			if err != nil {
				t.Error(err)
			}
			if votes.Count == 0 {
				t.Error("Expected at least one topic option vote")
			}
			for _, vote := range votes.Results {
				if vote.ID == "" {
					t.Error("Expected topic option vote ID")
				}
				if vote.GovernorVotingSession.Governor.Account.Address == "" {
					t.Error("Expected topic option vote governor address")
				}
				if vote.GovernorVotingSession.Transaction.ID == "" {
					t.Error("Expected topic option vote transaction ID")
				}
				if vote.GovernorVotingSession.Governor.CommittedAlgoAmount < 1 {
					t.Errorf("Expected topic option vote committed ALGO amount to be greater than 0, got %d", vote.GovernorVotingSession.Governor.CommittedAlgoAmount)
				}
				if vote.AllocatedAlgoAmount < 1 {
					t.Errorf("Expected topic option vote allocated ALGO amount to be greater than 0, got %d", vote.AllocatedAlgoAmount)
				}
				//lint:ignore SA4004 Only getting the first result
				break
			}
			//lint:ignore SA4004 Only getting the first result
			break
		}
		//lint:ignore SA4004 Only getting the first result
		break
	}
}

func TestGetTransaction(t *testing.T) {
	slug := "governance-period-2"
	address := "3RYOY2LTPC6GLT3ZYE4LUFGGAEMY7GRENZQO7RFNGK2LGCV77QNASK6C6Y"
	governorActivities, err := GetGovernorActivities(slug, address, "", "")
	if err != nil {
		t.Error(err)
	}
	for _, activity := range governorActivities.Results {
		transaction, err := GetTransaction(activity.Transaction.TransactionID)
		if err != nil {
			t.Error(err)
		}
		if transaction.ID != activity.Transaction.ID {
			t.Errorf("Expected transaction ID to be %s, got %s", activity.Transaction.ID, transaction.ID)
		}
		if transaction.GovernorActivity.ID != activity.ID {
			t.Errorf("Expected transaction governor activity ID to be %s, got %s", activity.ID, transaction.GovernorActivity.ID)
		}
		if transaction.GovernorActivity.Governor.Account.Address != address {
			t.Errorf("Expected transaction governor address to be %s, got %s", address, transaction.GovernorActivity.Governor.Account.Address)
		}
		//lint:ignore SA4004 Only getting the first result
		break
	}
}

func TestGetVotingSession(t *testing.T) {
	slug := "voting-session-q1-2022"
	votingSession, err := GetVotingSession(slug)
	if err != nil {
		t.Error(err)
	}
	if votingSession.ID != 5 {
		t.Errorf("Expected voting session ID to be 5, got %d", votingSession.ID)
	}
	if votingSession.Slug != slug {
		t.Errorf("Expected voting session slug to be %s, got %s", slug, votingSession.Slug)
	}
	foundTopic5 := false
	for _, topic := range votingSession.Topics {
		if topic.ID == 5 {
			foundTopic5 = true
		}
	}
	if !foundTopic5 {
		t.Errorf("Expected voting session to have topic with ID 5")
	}
	if !votingSession.HasFoundationProposal {
		t.Errorf("Expected voting session to have foundation proposal")
	}
}
