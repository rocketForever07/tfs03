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

	router := mux.NewRouter().StrictSlash(true)//StrictSlash :localhost/home/ -> localhost/home

	r := router.PathPrefix("/people").Subrouter()

	initHandlers(router)
	initHandlersPeople(r)

	log.Fatal(http.ListenAndServe(":10000", router))
	log.Println("Starting listening server on port 10000")


}

func initDB() {
	config := database.Config{

		ServerName: "localhost:3306",
		Username:   "root",
		Password:   "",
		DB:         "db7",
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
	// router.HandleFunc("/peoples", controller.GetAllPerson).Methods("GET")
	// router.HandleFunc("/people/{id}", controller.GetDetail).Methods("GET")
	// router.HandleFunc("/people", controller.CreatePerson).Methods("POST")
	// router.HandleFunc("/people/{id}", controller.UpdatePerson).Methods("PUT")
	// router.HandleFunc("/people/{id}", controller.DeletePerson).Methods("DELETE")

    router.Use(middleware.ContentTypeChecking)
	// router.Use(middleware.LoginChecking)

}

func initHandlersPeople(router *mux.Router){

	router.HandleFunc("", controller.GetAllPerson).Methods("GET")
	router.HandleFunc("/{id}", controller.GetDetail).Methods("GET")
	router.HandleFunc("", controller.CreatePerson).Methods("POST")
	router.HandleFunc("/{id}", controller.UpdatePerson).Methods("PUT")
	router.HandleFunc("/{id}", controller.DeletePerson).Methods("DELETE")

}
