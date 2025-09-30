package payment

type Payment struct {
	Id int
	Name string
	Description string
	Cron string
	Url string
	Company *string
	Group *string
}

type Company struct {
	Id int
	Name string
	Description string
	Industry string   
	Website *string
	Location *string
}


type Paymentdb struct {
	Id int
	Name string
	Cron string
	Url string
	Companyid int
	PaymentGroup int
}

type PaymentGroupdb struct {
	Id int
	Name string
	Description string
}

type Companydb struct {
	Id int
	Name string
	Description string
	Industry string   
	Website string
	Location string
}	

