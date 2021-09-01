package entity



type Category struct{
	ID int `gorm:primaryKey`
	Name string `json:"name"`

	Products []Product `gorm:"foreignKey:CategoryId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Category) TableName() string {
    return "db_category"
}