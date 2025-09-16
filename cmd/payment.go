package cmd

import (

	"github.com/spf13/cobra"
)



var paymentcmd = &cobra.Command{
	Use: "payment",
	Aliases: []string{"pmt", "paymt"},
	Short: "Base payment object to use payctl",
	Long: "Payment is the base object to use payctl, which will refer to any kind of payment you want to  create and manage with the tool",
} 


func init()  {
	rootCmd.AddCommand(paymentcmd)
}
