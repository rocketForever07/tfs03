package entity

type CartDetail struct{

	ID int `gorm:"primaryKey;autoIncrement:true"`
	CartId int
	ProductId int
	Quantity int
	Price float64
}


func (CartDetail) TableName() string {
    return "db_cart_detail"
}