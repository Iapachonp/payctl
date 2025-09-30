package cmd

import (
	"fmt"
	"payctl/database"
	"payctl/payment"
	"payctl/tables"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var createCompanycmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"cr", "new", "apply"},
	Short:   "Create a company",
	Long:    "Create a new company object",
	Run: func(cmd *cobra.Command, args []string) {
		// Command logic here
		var cmp string
		fmt.Println("Creating new company Object")
		name, _ := cmd.Flags().GetString("name")
		industry, _ := cmd.Flags().GetString("industry")
		description, _ := cmd.Flags().GetString("description")
		switch {
			case cmd.Flags().Lookup("location").Changed && cmd.Flags().Lookup("website").Changed:
			    website, _ := cmd.Flags().GetString("website")
			    location, _ := cmd.Flags().GetString("location")
			    cmp = fmt.Sprintf(
				"INSERT INTO companies (name, industry, website, description, location) VALUES ('%s', '%s', '%s', '%s', '%s');",
				name, industry, website, description, location)
			case cmd.Flags().Lookup("website").Changed:
			    website, _ := cmd.Flags().GetString("website")
			    cmp = fmt.Sprintf(
				"INSERT INTO companies (name, industry, website,description ) VALUES ('%s', '%s', '%s', '%s');",
				name, industry, website, description)
			case cmd.Flags().Lookup("location").Changed:
			    location, _ := cmd.Flags().GetString("location")
			    cmp = fmt.Sprintf(
				"INSERT INTO companies (name, industry, description, location) VALUES ('%s', '%s', '%s', '%s');",
				name, industry, description, location)
			default:
			    cmp = fmt.Sprintf(
				"INSERT INTO companies (name, industry, description) VALUES ('%s', '%s', '%s');",
				name, industry, description)
		}
		fmt.Println(cmp)
		db := database.Open()
		companyexec, err := db.Exec(cmp)
		if err != nil {
			fmt.Printf("Error creating Company: %v", err)
		}
		companyid, _ := companyexec.LastInsertId()
		fmt.Println(companyid)
		cmpy, err := payment.GetCompany(int(companyid))
		if err != nil {
			fmt.Printf("Error Fetching new company: %v", err)
		}
		tableRowHeader := table.Row{"id", "Name","Industry", "Description", "Website", "Location"}
		tableCaption := "New company created"
		payList := []table.Row{}
		payList = append(payList, table.Row{cmpy.Id, cmpy.Name, cmpy.Industry, cmpy.Description ,*cmpy.Website, *cmpy.Location })
		tables.PrintTable(tableRowHeader, tableCaption, payList)
	},
}

func init() {
	companycmd.AddCommand(createCompanycmd)
	createCompanycmd.PersistentFlags().String("name", "", "Name of the company object to create.")
	createCompanycmd.MarkFlagRequired("name")
	createCompanycmd.PersistentFlags().String("industry", "", "Name of the industry")
	createCompanycmd.MarkFlagRequired("industry")
	createCompanycmd.PersistentFlags().String("website", "", "Name of the website")
	createCompanycmd.MarkFlagRequired("website")
	createCompanycmd.PersistentFlags().String("description", "", "Description of the company object to create.")
	createCompanycmd.MarkFlagRequired("description")
	createCompanycmd.PersistentFlags().String("location", "", "Name of the company location of the company")
	createCompanycmd.MarkFlagRequired("location")
}
