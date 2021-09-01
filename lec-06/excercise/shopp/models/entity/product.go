package entity

type Product struct{
	ID 			int `gorm:primaryKey`
	Name 		string
	Price 		float64
	Quantity 	int
	Desciption 	string
	CategoryId 	int
	Status 		string
	OrderDetails [] OrderDetail `gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CartDetails [] CartDetail 	`gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}


func (Product) TableName() string {
    return "db_product"
}


