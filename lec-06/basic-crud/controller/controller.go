package controller

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"basic-crud/database"
	"net/http"

	"basic-crud/entity"

	"github.com/gorilla/mux"
)

func GetAllPerson(w http.ResponseWriter, r *http.Request) {

	//used to contain list person
	var persons []entity.Person

	database.Connector.Find(&persons)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(persons)

}

func GetDetail(w http.ResponseWriter, r *http.Request) {

	var person entity.Person

	Vars := mux.Vars(r)
	key := Vars["id"]

	database.Connector.First(&person, key)
	w.Header().Set("Content-type", "application/json")

	//NewEncoder(w):returns a new encoder that writes to w.
	//Encode(person):writes the JSON encoding of v
	json.NewEncoder(w).Encode(person)

}

func CreatePerson(w http.ResponseWriter, r *http.Request) {

	//read data from body request
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var person entity.Person

	//convert json to obj
	json.Unmarshal(reqBody, &person)

	log.Println("---------person\n", person, "\n-------end")

	//import to db
	database.Connector.Create(person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)

}

func UpdatePerson(w http.ResponseWriter, r *http.Request){

	// //read data from body request
	// reqBody, err := ioutil.ReadAll(r.Body)
	// if err != nil{
	// 	log.Println(err)
	// }

	// var person entity.Person
	// //convert json to obj
	// json.Unmarshal(reqBody,&person)

	// //update to db
	// database.Connector.Save(person)

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(person)

	requestBody, _ := ioutil.ReadAll(r.Body)
	var person entity.Person
	json.Unmarshal(requestBody, &person)

	log.Println("---------Reqbody\n", requestBody,"\n-------end")
	log.Println("---------person\n", person,"\n-------end")

	database.Connector.Save(&person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)

}


func DeletePerson(w http.ResponseWriter, r *http.Request) {


	vars := mux.Vars(r)

	key := vars["id"]

	var person entity.Person

	//convert to integer(tránh bị injection)
	id, _ := strconv.ParseInt(key, 10, 64)

	database.Connector.Where("id = ?", id).Delete(&person)
	w.WriteHeader(http.StatusNoContent)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	// fmt.Println("GET params were:", r.URL.Query())

	// username := r.URL.Query().Get("username")
	// password := r.URL.Query().Get("password")

	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println(username, password)
	w.Write([]byte("\ncheck authorization\n"))
	w.Write([]byte(`{"hello": "world"}`))

	if username == "khoidt" && password == "123abc"{
		token := "8282737kjgg32hh"
		w.Header().Set("Authorization",token)

		w.Write([]byte("authorized"))


	}else{
		w.Write([]byte("unauthorization"))
	}

	return

}
