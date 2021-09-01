package entity

import "time"

type Cart struct{
	ID int `gorm:"primaryKey;autoIncrement:true"`
	UserId int
	CreatedAt time.Time
	Status int
	Cost float64

	CartDetails [] CartDetail `gorm:"foreignKey:CartId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

}

func (Cart) TableName() string {
    return "db_cart"
}