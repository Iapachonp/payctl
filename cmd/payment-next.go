package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"payctl/payment"
	"payctl/tables"
	"github.com/jedib0t/go-pretty/v6/table"
)



var nextPaymentcmd = &cobra.Command{
	Use: "next",
	Aliases: []string{"nxt", "n", "nt"},
	Short: "get next payment date from a payment",
	Long: "next subcommand allows you to get the next payment date of a specic payment.",
	Run: func(cmd *cobra.Command, args []string) {
		// Command logic here
		fmt.Println("getting next payment date" )
		paymentId, _ := cmd.Flags().GetInt("id")
		deep, _ := cmd.Flags().GetInt("number")
		pmto, err := payment.GetPayment(paymentId)
		nextPayments , err := payment.GetNextpayments(paymentId, deep)
		if err != nil {
			fmt.Printf("Error Fetching next Payment dates: %v", err)
		}
		tableRowHeader := table.Row{"id", "Name", "Description", "Status", "Cron", "Url", "Company", "Group"}
		payList := []table.Row{}
		var status string
		if pmto.Status { 
			status = "enabled" 
		} else { 
			status = "disabled"
		}
		payList = append(payList, table.Row{pmto.Id, pmto.Name, pmto.Description, status, pmto.Cron, pmto.Url, *pmto.Company, *pmto.Group})
		tables.PrintTable(tableRowHeader, "", payList)
		// DATES table 
		tableRowHeader = table.Row{"#","Date"}
		tableCaption := "Next payment dates"
		dateList := []table.Row{}
		for i, date := range nextPayments { 
			if i != 0 {
				dateList = append(dateList, table.Row{i, date.Format("2006-01-02 15:04:00")})
			}
		}
		tables.PrintTable(tableRowHeader, tableCaption, dateList)
	},
} 


func init()  {
	paymentcmd.AddCommand(nextPaymentcmd)
	nextPaymentcmd.PersistentFlags().Int("id", 0, "Payment ID to get next payment date.")
	nextPaymentcmd.PersistentFlags().Int("number", 1, "Number of next payments to get.")
}
