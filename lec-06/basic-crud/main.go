package main

import (
	"basic-crud/controller"
	"basic-crud/database"
	"basic-crud/entity"
    "basic-crud/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	initDB()
	log.Println("Starting listening server on port 10000")

	router := mux.NewRouter().StrictSlash(true)
	initHandlers(router)

	log.Fatal(http.ListenAndServe(":10000", router))

}

func initDB() {
	config := database.Config{

		ServerName: "localhost:3306",
		Username:   "root",
		Password:   "",
		DB:         "learning",
	}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}

	database.Migrate(&entity.Person{})

}

func initHandlers(router *mux.Router) {

	router.HandleFunc("/home",func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("\nhome page\n"))
	})
	router.HandleFunc("/login",controller.LoginHandler)
	router.HandleFunc("/peoples", controller.GetAllPerson).Methods("GET")
	router.HandleFunc("/people/{id}", controller.GetDetail).Methods("GET")
	router.HandleFunc("/people", controller.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", controller.UpdatePerson).Methods("PUT")
	router.HandleFunc("/people/{id}", controller.DeletePerson).Methods("DELETE")

    router.Use(middleware.ContentTypeChecking)
	// router.Use(middleware.LoginChecking)

}
