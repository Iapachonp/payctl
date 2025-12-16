package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"payctl/payment"
	"payctl/tables"
	"github.com/jedib0t/go-pretty/v6/table"
)



var todaycmd = &cobra.Command{
	Use: "today",
	Aliases: []string{"td", "tdy", "now"},
	Short: "today helps you to get today's payments",
	Long: "today is the base object to get payments that needs to be paid today date.",
	Run: func(cmd *cobra.Command, args []string) {
		// Command logic here
		fmt.Println("getting payments for today" )
		payments, nextPayments , err := payment.GetTodaypayments()
		if err != nil  { fmt.Printf("Error processing today Payments: %v", err)}
		tableRowHeader := table.Row{"id", "Name", "Cron", "Url", "Company", "Next Payment Date"}
		tableCaption   := "today payments list"
		payList := []table.Row{}
		for i, pmt := range payments {
			if pmt.Status { 
				payList = append(payList, table.Row{pmt.Id, pmt.Name, pmt.Cron, pmt.Url, *pmt.Company, nextPayments[i].Format("2006-01-02 15:04:00")})
			} 
		}
		tables.PrintTable(tableRowHeader, tableCaption, payList)	
	},
} 


func init()  {
	rootCmd.AddCommand(todaycmd)
}
