package cmd

import (
	"fmt"
	"payctl/payment"
	"payctl/tables"
	"github.com/spf13/cobra"
	"github.com/jedib0t/go-pretty/v6/table"
)



var listPaymentscmd = &cobra.Command{
	Use: "list",
	Aliases: []string{"ps", "ls"},
	Short: "List payments",
	Long: "List available payments",
	Run: func(cmd *cobra.Command, args []string) {
		// Command logic here
		fmt.Println("listing available payments" )
		var limit int
		switch {
			case cmd.Flags().Lookup("limit").Changed:
				limit, _ = cmd.Flags().GetInt("limit")
			default:
				limit = 0
		}
		payments, err := payment.GetPayments(limit)
		if err != nil  { fmt.Printf("Error fetching Payments: %v", err)}
		tableRowHeader := table.Row{"id", "Name", "Description", "Cron", "Url", "Company", "Group"}
		tableCaption   := "payments list"
		payList := []table.Row{}
		for _, pmt := range payments {
			payList = append(payList, table.Row{pmt.Id, pmt.Name, pmt.Description, pmt.Cron, pmt.Url, *pmt.Company, *pmt.Group})
		}
		tables.PrintTable(tableRowHeader, tableCaption, payList)
	},
} 


func init()  {
	paymentcmd.AddCommand(listPaymentscmd)
	listPaymentscmd.PersistentFlags().Int("limit", 0, "Limit the number of the list of payments.")
}
