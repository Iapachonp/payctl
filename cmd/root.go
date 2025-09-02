/*
Copyright © 2025 sir-pachis

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "payctl",
	Short: "payctl is cli tool to manage your payments, taxes, services, credit cards, credits, mortgage, etc.",
	Long: `payctl is cli tool to manage your payments, taxes, services, credit cards, credits, mortgage, etc.
	it helps you with cron notifications, financial balance, and it can be integrated with tmux, and neovim, it has multiple pluggings
	so you can display important information, from comming payments overdue payments and state of your debts`,
	Run: func(cmd *cobra.Command, args []string) {
		println("this is my new payctl to pay my cobras")
	},	
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.payctl.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


