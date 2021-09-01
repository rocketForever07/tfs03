package main

import (
	"log"
	db "shopp/database"

	category "shopp/repository"
	"shopp/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main(){

	db.InitDB()
	routes.Init()
	log.Println("-------")


	log.Println(category.DeleteCategory(3))

}





