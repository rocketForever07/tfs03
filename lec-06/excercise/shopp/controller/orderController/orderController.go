package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"shopp/models/entity"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	repo "shopp/repository"
)

func GetAll(w http.ResponseWriter, r *http.Request){

	orders,err := repo.GetAllOrder()

	if err != nil{
		w.WriteHeader(http.StatusNoContent)
		log.Println(err)
	}else{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	
		json.NewEncoder(w).Encode(orders)
	}

}


func GetDetail(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	key := vars["id"]
	id ,_ := strconv.ParseInt(key,10,64)

	order,err := repo.GetOrderById(id)
	
	fmt.Println(order)
	if err != nil{
		w.WriteHeader(http.StatusNoContent)
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)

}


func Create(w http.ResponseWriter, r *http.Request){

	reqBody,err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Println(err)
	}

	var order entity.Order
	//convert json to obj
	json.Unmarshal(reqBody,&order)
	fmt.Println(order)

	//import to db
	newOrder,err := repo.CreateOrder(order)

	if err != nil{
		log.Println(err)	
		w.WriteHeader(http.StatusNoContent)
	}
	//
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	//response this payment obj(json)
	json.NewEncoder(w).Encode(newOrder)

}

func Update(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id,_ := strconv.Atoi(vars["id"])

	requestBody, _ := ioutil.ReadAll(r.Body)
	var order entity.Order
	json.Unmarshal(requestBody, &order)
	order.ID = id
	log.Println(order)

	updateOrder,err := repo.UpdateOrder(order)
	if err != nil{
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updateOrder)

}

func Delete(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	key := vars["id"]

	id ,_ := strconv.ParseInt(key,10,64)

	err := repo.DeleteOrder(id)
	if err!=nil{
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
}