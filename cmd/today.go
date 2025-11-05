package cmd

import (

	"github.com/spf13/cobra"
)



var todaycmd = &cobra.Command{
	Use: "today",
	Aliases: []string{"td", "tdy", "now"},
	Short: "Base today object to use payctl",
	Long: "today is the base object to get payments that needs to be paid today date.",
} 


func init()  {
	rootCmd.AddCommand(todaycmd)
}
