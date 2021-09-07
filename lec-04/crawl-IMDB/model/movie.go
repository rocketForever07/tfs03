package model

type Movie struct {

	ID int `gorm:"primaryKey;autoIncrement:true"`
	Name string 
	Year string
	Rate string
	ImgSrc string

}