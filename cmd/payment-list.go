package cmd

import (
	"fmt"
	"payctl/database"
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
		db := database.Open()
		list := "select p.id, p.Name, p.Cron, p.Url, c.name, g.name from payments p join companies c on p.companyid = c.id join paymentgroup g on p.paymentgroupid = g.id"
		payments, err := db.Query(list)
		if err != nil  { fmt.Printf("Error fetching Payments: %v", err)}
		defer payments.Close()
		tableRowHeader := table.Row{"id", "Name", "Cron", "Url", "Company", "Group"}
		tableCaption   := "payments list"
		payList := []table.Row{}
		for payments.Next() {
			var pmt payment.Payment 
			err := payments.Scan(&pmt.Id, &pmt.Name, &pmt.Cron, &pmt.Url, &pmt.Company, &pmt.Group)
			if err != nil  { fmt.Printf("Error unmarshall Payment: %v", err)}
			payList = append(payList, table.Row{pmt.Id, pmt.Name, pmt.Cron, pmt.Url, pmt.Company, pmt.Group})
		}
		tables.PrintTable(tableRowHeader, tableCaption, payList)
	},
} 


func init()  {
	paymentcmd.AddCommand(listPaymentscmd)
}
