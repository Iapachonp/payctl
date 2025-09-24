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


func GetPayment(paymentId int)(Payment, error){
	db := database.Open()
	var payment Payment 
	query := "select p.id, p.Name, p.Cron, p.Url, c.name, g.name from payments p join companies c on p.companyid = c.id join paymentgroup g on p.paymentgroupid = g.id where p.id = ?"
	err := db.QueryRow(query, paymentId).Scan(&payment.Id, &payment.Name, &payment.Cron, &payment.Url, &payment.Company, &payment.Group)
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
