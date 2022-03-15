package rewards

import (
	"testing"

	"github.com/Dan6erbond/algogovernance/pkg/client"
	"github.com/Dan6erbond/algogovernance/pkg/constants"
	"github.com/Dan6erbond/algogovernance/pkg/utils"
)

func TestGetGovernorRewardsForPeriod(t *testing.T) {
	period := client.Period{
		Slug:                   "governance-period-1",
		TotalCommittedStake:    1000 / constants.MICRO_ALGO,
		AlgoAmountInRewardPool: int(100.0 / constants.MICRO_ALGO),
	}
	governor := client.Governor{
		CommittedAlgoAmount: int(100.0 / constants.MICRO_ALGO),
	}

	result := GetGovernorRewardsForPeriod(period, &governor)

	if utils.MicroAlgoToAlgo(result) != 10.0 {
		t.Errorf("Expected 10 ALGO, got %f ALGO", result)
	}
}
