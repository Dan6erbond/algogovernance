package pkg

import (
	"github.com/Dan6erbond/algogovernance/internal/algogovernance"
)

func GetRewards(periodSlug string, governorAddress string) (result float64, err error) {
	period, err := algogovernance.GetPeriod(periodSlug)

	if err != nil {
		return 0, err
	}

	return GetRewardsForPeriod(period, governorAddress)
}

func GetRewardsForCurrentPeriod(governorAddress string) (result float64, err error) {
	period, err := algogovernance.GetActivePeriod()

	if err != nil {
		return 0, err
	}

	return GetRewardsForPeriod(period, governorAddress)
}

func GetRewardsForPeriod(period algogovernance.GovernancePeriod, governorAddress string) (result float64, err error) {
	governor, err := algogovernance.GetGovernors(period.Slug, governorAddress)

	if err != nil {
		return 0, err
	}
	result = float64(governor.Governor.CommittedAlgoAmount) / period.TotalCommittedStake * float64(period.AlgoAmountInRewardPool)
	result = result * MICRO_ALGO

	return result, err
}
