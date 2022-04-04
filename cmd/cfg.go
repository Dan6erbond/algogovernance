package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cfgCmd represents the cfg command
var cfgCmd = &cobra.Command{
	Use:   "cfg",
	Short: "View the configuration parameters detected by flags and Viper",
	Long:  `Use this command to ensure configuration variables are set correctly.`,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.ConfigFileUsed() != "" {
			fmt.Printf("Loaded config from %s:\n", viper.ConfigFileUsed())
		} else {
			home, err := os.UserHomeDir()
			cobra.CheckErr(err)
			fmt.Printf("No config file found. Create one at %s.\n", path.Join(home, "algogovernance.yaml"))
		}

		if governor != "" {
			fmt.Println("Governor:", governor)
		} else {
			if gov := viper.GetString("governor"); gov != "" {
				fmt.Println("Governor:", gov)
			} else {
				fmt.Println("No governor specified")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cfgCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cfgCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cfgCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
