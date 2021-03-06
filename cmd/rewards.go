package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/Dan6erbond/algogovernance/pkg/client"
	algoRewards "github.com/Dan6erbond/algogovernance/pkg/rewards"
	"github.com/Dan6erbond/algogovernance/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rewardsCmd represents the rewards command
var rewardsCmd = &cobra.Command{
	Use:   "rewards",
	Short: "Get the rewards for a selected governor and period",
	Long:  `Use this command to view the expected rewards for a governor's wallet address in a governance period.`,
	Run: func(cmd *cobra.Command, args []string) {
		period := cmd.Flag("period").Value.String()
		if governor == "" {
			governor = viper.GetString("governor")
		}
		if governor == "" {
			log.Fatalf("Governor address is required, try again by setting it in the .env file or using the -g flag.")
		}
		var (
			governancePeriod client.Period
			rewards          float64
			err              error
		)
		if period == "" {
			governancePeriod, err = client.GetActivePeriod()
		} else {
			if !strings.HasPrefix(period, "governance-period-") {
				period = "governance-period-" + period
			}
			governancePeriod, err = client.GetPeriod(period)
		}

		if err != nil {
			log.Fatalf("Error getting active period: %s", err)
		}

		rewards, err = algoRewards.GetRewardsForPeriod(governancePeriod, governor)

		if err != nil {
			log.Fatalf("Error getting active period: %s", err)
		}

		fmt.Printf("Expected rewards: %.2f ALGO\n", utils.MicroAlgoToAlgo(rewards))
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
