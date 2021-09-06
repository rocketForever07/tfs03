package storage

import(
	"gorm.io/gorm"
	"log"
	model "crawl-IMDB/model"
	"fmt"
	"gorm.io/driver/mysql"

)


var DB *gorm.DB
var err error
func InitDB(){

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", "root", "dtkhoi07", "localhost:3306", "db4")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database %s", err)
	}

	err := DB.AutoMigrate(&model.Movie{})
		
	if err != nil {
		log.Fatal("failed to load migrations")
	}

	log.Println("successful migrate")

}

func AddToDB(movie model.Movie)(model.Movie,error){
	result := DB.Create(&movie)

	if result.Error != nil{
		return model.Movie{},result.Error
	}
	return movie,nil
}
