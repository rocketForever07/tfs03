package entity

type Order struct{
	ID int `json:"id"`
	CustomerId int `json:"customerId"`
	Cost int `json:"cost"`
	Status int `json:"status"`
	PaymentId int `json:"paymentId"`
	Payment Payment `json:"-"`
}

