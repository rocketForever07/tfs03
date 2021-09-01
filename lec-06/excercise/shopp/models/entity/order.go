package entity

type Order struct{
	ID int `gorm:"primaryKey;autoIncrement:true"`
	UserId int 
	Cost float64 
	Status int

	OrderDetails []OrderDetail `gorm:"foreignKey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Order) TableName() string {
    return "db_order"
}

