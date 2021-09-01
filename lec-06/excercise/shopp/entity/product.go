package entity

type Product struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Quanlity int `json:"quanlity"`
	Price float64 `json:"price"`
	Desciption string `json:"desc"`
}


func (Product) TableName() string {
    return "db_product"
}
