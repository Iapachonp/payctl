package payment

import (
	"fmt"
	"payctl/database"
)

func GetGroup(groupId int) (PaymentGroupdb, error) {
	db := database.Open()
	var group PaymentGroupdb 
	query := "select * from paymentGroup where id = ?"
	err := db.QueryRow(query, groupId).Scan(&group.Id, &group.Name, &group.Description)
	if err != nil {
		return PaymentGroupdb{}, fmt.Errorf("error fetching group: %w", err)
	}
	return group, nil
}

func GetCompany(companyId int) (Companydb, error) {
	db := database.Open()
	var company Companydb 
	query := "select * from companies where id = ?"
	err := db.QueryRow(query, companyId).Scan(&company.Id, &company.Name, &company.Industry, &company.Website, &company.Location)
	if err != nil {
		return Companydb{}, fmt.Errorf("error fetching company: %w", err)
	}
	return company, nil
}

func GetCompanies(limit int) ([]Companydb, error) {
	db := database.Open()
	var companies []Companydb
	var query string
	if limit > 0 {query = "select * from companies limit ?" } else {query = "select * from companies"} 
	cmps, err := db.Query(query, limit)
	defer cmps.Close()
	if err != nil {
		return nil, fmt.Errorf("error fetching payments: %w", err)
	}
	for cmps.Next() {
		var company Companydb 
		err = cmps.Scan(&company.Id, &company.Name, &company.Description, &company.Industry, &company.Location, &company.Website)
		if err != nil {
			return nil, fmt.Errorf("Error unmarshall Company: %w", err)
		}
		companies = append(companies, company )
	}
	
	return companies, nil
}


func GetPayments(limit int)([]Payment, error){
	db := database.Open()
	var payments[] Payment
	var query string
	if limit > 0 {
		query = "select p.id, p.Name, p.Description, p.Cron, p.Url, c.name, g.name from payments p left join companies c on p.companyid = c.id left join paymentgroup g on p.paymentgroupid = g.id limit ?;" 
	} else {
		query = "select p.id, p.Name, p.Description, p.Cron, p.Url, c.name, g.name from payments p left join companies c on p.companyid = c.id left join paymentgroup g on p.paymentgroupid = g.id;"
	}
	pmts, err := db.Query(query, limit)
	defer pmts.Close()
	if err != nil {
		return nil, fmt.Errorf("error fetching payments: %w", err)
	}
	for pmts.Next() {
		var pmt Payment 
		err := pmts.Scan(&pmt.Id, &pmt.Name, &pmt.Description, &pmt.Cron, &pmt.Url, &pmt.Company, &pmt.Group)
		if err != nil  { fmt.Printf("Error unmarshall Payment: %v", err)}
		na := "n/a"
		if pmt.Company == nil {
			pmt.Company = &na
		}
		if pmt.Group == nil {
			pmt.Group = &na
		}
		payments = append(payments, pmt)
	}
	return payments, nil
}


func GetPayment(paymentId int)(Payment, error){
	db := database.Open()
	var payment Payment 
	query := "select p.id, p.Name, p.Cron, p.Url, c.name, g.name from payments p left join companies c on p.companyid = c.id left join paymentgroup g on p.paymentgroupid = g.id where p.id = ?"
	err := db.QueryRow(query, paymentId).Scan(&payment.Id, &payment.Name, &payment.Cron, &payment.Url, &payment.Company, &payment.Group)
	na := "n/a"
	if payment.Company == nil {
		payment.Company = &na
	}
	if payment.Group == nil {
		payment.Group = &na
	}
	if err != nil {
		return Payment{}, fmt.Errorf("error fetching payment: %w", err)
	}
	return payment, nil
}

func GetPaymentdb(paymentId int) (Paymentdb, error) {
	db := database.Open()
	var payment Paymentdb 
	query := "select * from payments where id = ?"
	err := db.QueryRow(query, paymentId).Scan(&payment.Id, &payment.Name, &payment.Cron, &payment.Url, &payment.Companyid, &payment.PaymentGroup)
	if err != nil {
		return Paymentdb{}, fmt.Errorf("error fetching payment: %w", err)
	}
	return payment, nil
}
