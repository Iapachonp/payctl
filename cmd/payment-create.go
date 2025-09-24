package cmd

import (
	"fmt"
	"payctl/database"
	"payctl/payment"
	"payctl/tables"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var createPaymentscmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"cr", "new", "apply"},
	Short:   "Create a payment",
	Long:    "Create a new payment object",
	Run: func(cmd *cobra.Command, args []string) {
		// Command logic here
		var pmt string
		fmt.Println("Creating new payment Object")
		name, _ := cmd.Flags().GetString("name")
		cron, _ := cmd.Flags().GetString("cron")
		url, _ := cmd.Flags().GetString("url")
		switch {
			case cmd.Flags().Lookup("companyid").Changed && cmd.Flags().Lookup("groupid").Changed:
			    companyid, _ := cmd.Flags().GetInt("companyid")
			    groupid, _ := cmd.Flags().GetInt("groupid")
			    pmt = fmt.Sprintf(
				"INSERT INTO payments (name, cron, url, companyid, paymentgroupid) VALUES ('%s', '%s', '%s', '%d', '%d');",
				name, cron, url, companyid, groupid,)
			case cmd.Flags().Lookup("companyid").Changed:
			    companyid, _ := cmd.Flags().GetInt("companyid")
			    pmt = fmt.Sprintf(
				"INSERT INTO payments (name, cron, url, companyid) VALUES ('%s', '%s', '%s', '%d');",
				name, cron, url, companyid,)
			case cmd.Flags().Lookup("groupid").Changed:
			    paymentgroupid, _ := cmd.Flags().GetInt("groupid")
			    pmt = fmt.Sprintf(
				"INSERT INTO payments (name, cron, url, paymentgroupid) VALUES ('%s', '%s', '%s', '%d');",
				name, cron, url, paymentgroupid,)
			default:
			    pmt = fmt.Sprintf(
				"INSERT INTO payments (name, cron, url) VALUES ('%s', '%s', '%s');",
				name, cron, url,)
		}
		fmt.Println(pmt)
		db := database.Open()
		paymentexec, err := db.Exec(pmt)
		if err != nil {
			fmt.Printf("Error creating Payment: %v", err)
		}
		paymentid, _ := paymentexec.LastInsertId()
		pmto, err := payment.GetPayment(int(paymentid))
		if err != nil {
			fmt.Printf("Error Fetching new Payment: %v", err)
		}
		tableRowHeader := table.Row{"id", "Name", "Cron", "Url", "Company", "Group"}
		tableCaption := "New payment created"
		payList := []table.Row{}
		payList = append(payList, table.Row{pmto.Id, pmto.Name, pmto.Cron, pmto.Url, pmto.Company, pmto.Group})
		tables.PrintTable(tableRowHeader, tableCaption, payList)
	},
}

func init() {
	paymentcmd.AddCommand(createPaymentscmd)
	paymentcmd.PersistentFlags().String("name", "", "Name of the payment object to create.")
	paymentcmd.MarkFlagRequired("name")
	paymentcmd.PersistentFlags().String("cron", "", "Name of the payment object to create.")
	paymentcmd.MarkFlagRequired("cron")
	paymentcmd.PersistentFlags().String("url", "", "Name of the payment object to create.")
	paymentcmd.MarkFlagRequired("url")
	// paymentcmd.PersistentFlags().String("group", "", "Name of the payment object to create.")
	// paymentcmd.PersistentFlags().String("company", "", "Name of the payment object to create.")
	paymentcmd.PersistentFlags().Int("groupid", 0, "Name of the payment object to create.")
	paymentcmd.PersistentFlags().Int("companyid", 0, "Name of the payment object to create.")
}
