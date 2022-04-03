// Package rewards provides utilities for calculating rewards in Algorand's Governance.
package rewards

import (
	"github.com/Dan6erbond/algogovernance/pkg/client"
)

// GetRewards calculates the rewards for a given period and governor using the period slug and governor address.
func GetRewards(periodSlug string, governorAddress string) (result float64, err error) {
	period, err := client.GetPeriod(periodSlug)

	if err != nil {
		return 0, err
	}

	return GetRewardsForPeriod(period, governorAddress)
}

// GetRewardsForCurrentPeriod calculates the rewards for the current active period by providing the governor's address.
func GetRewardsForCurrentPeriod(governorAddress string) (result float64, err error) {
	period, err := client.GetActivePeriod()

	if err != nil {
		return 0, err
	}

	return GetRewardsForPeriod(period, governorAddress)
}

// GetRewardsForPeriod calculates the rewards for a client.Period object and fetches the governor by their wallet address.
func GetRewardsForPeriod(period client.Period, governorAddress string) (result float64, err error) {
	governor, err := client.GetPeriodGovernorStatus(period.Slug, governorAddress)

	if err != nil {
		return 0, err
	}

	return GetGovernorRewardsForPeriod(period, &governor.Governor), nil
}

// GetGovernorRewardsForPeriod calculates the rewards for a client.Period object and a client.Governor object.
func GetGovernorRewardsForPeriod(period client.Period, governor client.AlgoCommitter) float64 {
	result := governor.GetCommittedAlgoAmount() / period.TotalCommittedStake * float64(period.AlgoAmountInRewardPool)
	return result
}
