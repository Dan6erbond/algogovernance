/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/Dan6erbond/algogovernance/internal/algogovernance"
	"github.com/Dan6erbond/algogovernance/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rewardsCmd represents the rewards command
var rewardsCmd = &cobra.Command{
	Use:   "rewards",
	Short: "Get the rewards for a selected governor and period",
	Long: `Use this command to view the expected rewards for a governor's wallet address in a governance period.`,
	Run: func(cmd *cobra.Command, args []string) {
		period := cmd.Flag("period").Value.String()
		if governor == "" {
			governor = viper.GetString("governor")
		}
		if governor == "" {
			log.Fatalf("Governor address is required, try again by setting it in the .env file or using the -g flag.")
		}
		var (
			governancePeriod algogovernance.GovernancePeriod
			rewards          float64
			err              error
		)
		if period == "" {
			governancePeriod, err = algogovernance.GetActivePeriod()
		} else {
			if !strings.HasPrefix(period, "governance-period-") {
				period = "governance-period-" + period
			}
			governancePeriod, err = algogovernance.GetPeriod(period)
		}

		if err != nil {
			log.Fatalf("Error getting active period: %s", err)
		}

		rewards, err = pkg.GetRewardsForPeriod(governancePeriod, governor)

		if err != nil {
			log.Fatalf("Error getting active period: %s", err)
		}

		fmt.Printf("Expected rewards: %.5f ALGO\n", rewards)
	},
}

func init() {
	rootCmd.AddCommand(rewardsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rewardsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rewardsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rewardsCmd.Flags().StringP("period", "p", "", "Governance period slug or ID")
}
