package database

import(
	"gorm.io/gorm"
	"log"
	entity "shopp/models/entity"
	"fmt"
	"gorm.io/driver/mysql"

)


var DB *gorm.DB
var err error
func InitDB(){

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", "root", "dtkhoi07", "localhost:3306", "db11")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database %s", err)
	}

	DB.AutoMigrate(&entity.User{},
		&entity.Category{},  
		&entity.Product{}, 
		&entity.Payment{}, 
		&entity.Order{}, 
		&entity.OrderDetail{}, 
		&entity.Cart{}, 
		&entity.CartDetail{})
	if err != nil {
		log.Fatal("failed to load migrations")
	}

	log.Println("successful migrate")

}
