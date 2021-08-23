package database

import(
	"basic-crud/entity"
	"github.com/jinzhu/gorm"
	"log"
)


//used for crud
var Connector *gorm.DB


//connect mysql using gorm
func Connect(connectString string) error{

	var err error
	Connector, err = gorm.Open("mysql",connectString)

	if err != nil{
		return err
	}

	log.Println("Connection successfully")
	return nil

}


//Migrate create/updates database table
func Migrate(table *entity.Person) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}
