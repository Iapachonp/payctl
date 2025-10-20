package cmd

import (
	"fmt"
	"payctl/payment"
	"payctl/tables"
	"github.com/spf13/cobra"
	"github.com/jedib0t/go-pretty/v6/table"
)



var listGroupscmd = &cobra.Command{
	Use: "list",
	Aliases: []string{"ps", "ls"},
	Short: "List groups",
	Long: "List available groups",
	Run: func(cmd *cobra.Command, args []string) {
		// Command logic here
		fmt.Println("listing available groups" )
		var limit int
		switch {
			case cmd.Flags().Lookup("limit").Changed:
				limit, _ = cmd.Flags().GetInt("limit")
			default:
				limit = 0
		}
		groups, err := payment.GetGroups(limit)
		if err != nil  { fmt.Printf("Error fetching Groups: %v", err)}
		tableRowHeader := table.Row{"id", "Name", "Description", "Cron", "Url", "Company", "Group"}
		tableCaption   := "groups list"
		payList := []table.Row{}
		for _, pmt := range groups {
			payList = append(payList, table.Row{pmt.Id, pmt.Name, pmt.Description, pmt.Cron, pmt.Url, *pmt.Company, *pmt.Group})
		}
		tables.PrintTable(tableRowHeader, tableCaption, payList)
	},
} 


func init()  {
	paymentcmd.AddCommand(listGroupscmd)
	listGroupscmd.PersistentFlags().Int("limit", 0, "Limit the number of the list of groups.")
}
