package cmd

import (
	"fmt"
	"go/printer"
	"payctl/database"
	"payctl/payment"
	"payctl/tables"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)



var createPaymentscmd = &cobra.Command{
	Use: "create",
	Aliases: []string{"cr", "new", "apply"},
	Short: "Create a payment",
	Long: "Create a new payment object",
	Run: func(cmd *cobra.Command, args []string) {
		// Command logic here
		var pmt string
		fmt.Println("Creating new payment Object" )
		db := database.Open()
		name, _ := cmd.Flags().GetString("name")
		cron, _ := cmd.Flags().GetString("cron")
		url, _ := cmd.Flags().GetString("url")
		if  ( cmd.Flags().Lookup("companyid").Changed ) && ( cmd.Flags().Lookup("groupid").Changed) {	
			companyid, _ := cmd.Flags().GetString("companyid")
			groupid, _ := cmd.Flags().GetString("groupid")
			pmt = fmt.Sprintf("INSERT INTO payments (name, cron, url, companyid, groupid) VALUES (%s, %s, %s, %n, %n);", name, cron, url, companyid, groupid )
		} else {
			if  ( cmd.Flags().Lookup("companyid").Changed ) {
				companyid, _ := cmd.Flags().GetString("companyid")
				pmt = fmt.Sprintf("INSERT INTO payments (name, cron, url, companyid, groupid) VALUES (%s, %s, %s, %n);", name, cron, url, companyid)
			} else {
				if  ( cmd.Flags().Lookup("groupid").Changed ) {
					groupid, _ := cmd.Flags().GetString("groupid")
					pmt = fmt.Sprintf("INSERT INTO payments (name, cron, url, companyid, groupid) VALUES (%s, %s, %s, %n);", name, cron, url, groupid)
				} else {
					pmt = fmt.Sprintf("INSERT INTO payments (name, cron, url, companyid, groupid) VALUES (%s, %s, %s);", name, cron, url)	
				} 			
			}
		}
		
		paymentexec, err := db.Exec(pmt)
		if err != nil  { fmt.Printf("Error creating Payment: %v", err)}
		paymentid, _ := paymentexec.LastInsertId()
		query := fmt.Sprintf("select p.id, p.Name, p.Cron, p.Url, c.name, g.name from payments p join companies c on p.companyid = c.id join paymentgroup g on p.paymentgroupid = g.id where p.id = %n", paymentid)
		payment, err := db.Query(query)
		if err != nil  { fmt.Printf("Error Fetching new created Payment: %v", err)}
		defer payment.Close()
		tableRowHeader := table.Row{"id", "Name", "Cron", "Url", "Company", "Group"}
		tableCaption   := "New payment created"
		payList := []table.Row{}
		for payment.Next() {
			var pmt payment.Payment 
			err := payment.Scan(&pmt.Id, &pmt.Name, &pmt.Cron, &pmt.Url, &pmt.Company, &pmt.Group)
			if err != nil  { fmt.Printf("Error unmarshall Payment: %v", err)}
			payList = append(payList, table.Row{pmt.Id, pmt.Name, pmt.Cron, pmt.Url, pmt.Company, pmt.Group})
		}
		tables.PrintTable(tableRowHeader, tableCaption, payList)
	},
} 


func init()  {
	paymentcmd.AddCommand(listPaymentscmd)
	paymentcmd.PersistentFlags().String("name", "", "Name of the payment object to create.")
	paymentcmd.MarkFlagRequired("name")
	paymentcmd.PersistentFlags().String("cron", "", "Name of the payment object to create.")
	paymentcmd.MarkFlagRequired("cron")
	paymentcmd.PersistentFlags().String("url", "", "Name of the payment object to create.")
	paymentcmd.MarkFlagRequired("url")
	// paymentcmd.PersistentFlags().String("group", "", "Name of the payment object to create.")
	// paymentcmd.PersistentFlags().String("company", "", "Name of the payment object to create.")
	paymentcmd.PersistentFlags().Int("groupid", 0, "Name of the payment object to create.")
	paymentcmd.PersistentFlags().Int("companyid", 0 ,"Name of the payment object to create.")
}
