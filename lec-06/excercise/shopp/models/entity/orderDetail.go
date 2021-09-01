package entity

type OrderDetail struct {
	ID        int `gorm:"primaryKey;autoIncrement:true"`
	ProductId int 
	OrderId   int    
	Quantity  int	
	Price 	  float64
					  
}


func (OrderDetail) TableName() string {
    return "db_order_detail"
}