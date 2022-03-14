/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cfgCmd represents the cfg command
var cfgCmd = &cobra.Command{
	Use:   "cfg",
	Short: "View the configuration parameters detected by flags and Viper",
	Long: `Use this command to ensure configuration variables are set correctly.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cfg called")
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
