package cmd

import (
	"fmt"
	"payctl/database"
	"payctl/payment"
	"payctl/tables"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var createGroupcmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"cr", "new", "apply"},
	Short:   "Create a payment",
	Long:    "Create a new payment object",
	Run: func(cmd *cobra.Command, args []string) {
		// Command logic here
		var grp string
		fmt.Println("Creating new group Object")
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		switch {
			default:
			    grp = fmt.Sprintf(
				"INSERT INTO paymentGroup (name,description) VALUES ('%s', '%s');",
				name, description)
		}
		fmt.Println(grp)
		db := database.Open()
		groupexec, err := db.Exec(grp)
		if err != nil {
			fmt.Printf("Error creating Group: %v", err)
		}
		groupid, _ := groupexec.LastInsertId()
		fmt.Println(groupid)
		grpo, err := payment.GetGroup(int(groupid))
		if err != nil {
			fmt.Printf("Error Fetching new Group: %v", err)
		}
		tableRowHeader := table.Row{"id", "Name", "Description"}
		tableCaption := "New group created"
		payList := []table.Row{}
		payList = append(payList, table.Row{grpo.Id, grpo.Name, grpo.Description})
		tables.PrintTable(tableRowHeader, tableCaption, payList)
	},
}

func init() {
	groupcmd.AddCommand(createGroupcmd)
	createGroupcmd.PersistentFlags().String("name", "", "Name of the group object to create.")
	createGroupcmd.MarkFlagRequired("name")
	createGroupcmd.PersistentFlags().String("description", "", "Description of the group object to create.")
	createGroupcmd.MarkFlagRequired("description")
}
