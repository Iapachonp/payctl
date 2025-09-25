package cmd

import (
	"fmt"
	"payctl/payment"
	"payctl/tables"
	"github.com/spf13/cobra"
	"github.com/jedib0t/go-pretty/v6/table"
)



var listCompaniescmd = &cobra.Command{
	Use: "list",
	Aliases: []string{"ps", "ls"},
	Short: "List companies",
	Long: "List available companies",
	Run: func(cmd *cobra.Command, args []string) {
		// Command logic here
		fmt.Println("listing available companies" )
		var companies []payment.Companydb
		var err error
		switch {
			case cmd.Flags().Lookup("limit").Changed:
				limit, _ := cmd.Flags().GetInt("limit")
				companies, err = payment.GetCompanies(limit)
				if err != nil  { fmt.Printf("Error fetching companies: %v", err)}
			default:
				companies, err = payment.GetCompanies(0)
				if err != nil  { fmt.Printf("Error fetching companies: %v", err)}
		}
		tableRowHeader := table.Row{"id", "Name", "Description","Industry", "Website", "Location"}
		tableCaption   := "companies list"
		comList := []table.Row{}
		for _, cmp := range companies {
			comList = append(comList, table.Row{cmp.Id, cmp.Name, cmp.Description, cmp.Industry, cmp.Website, cmp.Location})
		}
		tables.PrintTable(tableRowHeader, tableCaption, comList)
	},
} 


func init()  {
	companycmd.AddCommand(listCompaniescmd)
	companycmd.PersistentFlags().Int("limit", 0, "Limit the numbers of companies result.")
}
