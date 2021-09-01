package entity

import(
	"time"
)

type Payment struct{
	ID int `gorm:"primaryKey;autoIncrement:true"`
	OrderId int
	Cost float64
	PaymentDate time.Time

	Order Order `gorm:"foreignKey:OrderId"`
}


func (Payment) TableName() string {
    return "db_payment"
}
