package rewards

import (
	"github.com/Dan6erbond/algogovernance/pkg/client"
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

func GetRewardsForPeriod(period client.Period, governorAddress string) (result float64, err error) {
	governor, err := client.GetPeriodGovernorStatus(period.Slug, governorAddress)

	if err != nil {
		return 0, err
	}

	return GetGovernorRewardsForPeriod(period, &governor.Governor), nil
}

func GetGovernorRewardsForPeriod(period client.Period, governor client.AlgoCommitter) float64 {
	result := governor.GetCommittedAlgoAmount() / period.TotalCommittedStake * float64(period.AlgoAmountInRewardPool)
	return result
}
