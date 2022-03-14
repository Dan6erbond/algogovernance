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

// currentPeriodCmd represents the currentPeriod command
var currentPeriodCmd = &cobra.Command{
	Use:   "currentPeriod",
	Short: "Get an overview of the current governance period",
	Long: `See an overview of the current governance period showing total locked stake, registration end, ALGO reward pool.

	Also shows governance rewards if a governor is configured in configuration files or via the -g flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		activePeriod, err := algogovernance.GetActivePeriod()

		if err != nil {
			log.Fatalf("Error getting rewards for governor: %s %s", governor, err)
		}

		fmt.Printf("Governance period: %s\n", strings.TrimPrefix(activePeriod.Slug, "governance-period-"))
		fmt.Printf("Total locked stake: %.5f ALGO", activePeriod.TotalCommittedStake)
		fmt.Printf("\nRegistration end: %s\n", activePeriod.RegistrationEndDatetime)

		if governor == "" {
			governor = viper.GetString("governor")
		}
		if governor != "" {
			rewards, err := pkg.GetRewardsForPeriod(activePeriod, governor)
			if err != nil {
				log.Fatalf("Error getting rewards for governor: %s %s", governor, err)
			}
			fmt.Printf("Expected rewards: %.5f ALGO\n", rewards)
		}
	},
}

func init() {
	rootCmd.AddCommand(currentPeriodCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// currentPeriodCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// currentPeriodCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
