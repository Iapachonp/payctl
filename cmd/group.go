package cmd

import (

	"github.com/spf13/cobra"
)



var groupcmd = &cobra.Command{
	Use: "group",
	Aliases: []string{"gr", "grp", "gru"},
	Short: "Base group object to use payctl",
	Long: "Group is the base object to use payctl, which will refer to any kind of payment groups you want to  create and manage with the tool",
} 


func init()  {
	rootCmd.AddCommand(groupcmd)
}
