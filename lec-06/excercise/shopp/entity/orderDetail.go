package entity

type OrderDetail struct {
	ID        int     `json:"id"`
	ProductId int     `json:"productId"`
	OrderId   int     `json:"orderId"`
	Product   Product `json:"-" gorm:"foreignKey:ProductId"`
	Order     Order   `json:"-" gorm:"foreignKey:OrderId"`
					  
}


func (OrderDetail) TableName() string {
    return "db_order_detail"
}