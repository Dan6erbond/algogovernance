package rewards

import (
	"github.com/Dan6erbond/algogovernance/pkg/client"
	"github.com/Dan6erbond/algogovernance/pkg/constants"
)

func GetRewards(periodSlug string, governorAddress string) (result float64, err error) {
	period, err := client.GetPeriod(periodSlug)

	if err != nil {
		return 0, err
	}

	return GetRewardsForPeriod(period, governorAddress)
}

func GetRewardsForCurrentPeriod(governorAddress string) (result float64, err error) {
	period, err := client.GetActivePeriod()

	if err != nil {
		return 0, err
	}

	return GetRewardsForPeriod(period, governorAddress)
}

func GetRewardsForPeriod(period client.GovernancePeriod, governorAddress string) (result float64, err error) {
	governor, err := client.GetGovernors(period.Slug, governorAddress)

	if err != nil {
		return 0, err
	}
	result = float64(governor.Governor.CommittedAlgoAmount) / period.TotalCommittedStake * float64(period.AlgoAmountInRewardPool)
	result = result * constants.MICRO_ALGO

	return result, err
}
