package cmd

import (

	"github.com/spf13/cobra"
)



var companycmd = &cobra.Command{
	Use: "company",
	Aliases: []string{"cm", "cmp", "cpy"},
	Short: "Base company object to use payctl",
	Long: "Payment is the base object to use payctl, which will refer to any kind of payment you want to  create and manage with the tool",
} 


func init()  {
	rootCmd.AddCommand(companycmd)
}
