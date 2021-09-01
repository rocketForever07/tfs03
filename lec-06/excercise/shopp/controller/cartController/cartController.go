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

	carts,err := repo.GetAllCart()

	if err != nil{
		w.WriteHeader(http.StatusNoContent)
		log.Println(err)
	}else{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	
		json.NewEncoder(w).Encode(carts)
	}

}


func GetDetail(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	key := vars["id"]
	id ,_ := strconv.ParseInt(key,10,64)

	cart,err := repo.GetCartById(id)
	
	fmt.Println(cart)
	if err != nil{
		w.WriteHeader(http.StatusNoContent)
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cart)

}


func Create(w http.ResponseWriter, r *http.Request){

	reqBody,err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Println(err)
	}

	var cart entity.Cart
	//convert json to obj
	json.Unmarshal(reqBody,&cart)
	fmt.Println(cart)

	//import to db
	newCart,err := repo.CreateCart(cart)

	if err != nil{
		log.Println(err)	
		w.WriteHeader(http.StatusNoContent)
	}
	//
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	//response this payment obj(json)
	json.NewEncoder(w).Encode(newCart)

}


func Update(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id,_ := strconv.Atoi(vars["id"])

	requestBody, _ := ioutil.ReadAll(r.Body)
	var cart entity.Cart
	json.Unmarshal(requestBody, &cart)
	cart.ID = id
	log.Println(cart)

	updateCart,err := repo.UpdateCart(cart)
	if err != nil{
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updateCart)

}

func Delete(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	key := vars["id"]

	id ,_ := strconv.ParseInt(key,10,64)

	err := repo.DeleteCart(id)
	if err!=nil{
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
}

func GetByUserId(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userId ,_ :=  strconv.ParseInt(vars["userId"],10,64)

	carts,err := repo.FindCartByUserId(userId)

	if err != nil{
		log.Println(err)
	}

	json.NewEncoder(w).Encode(carts)
	w.WriteHeader(http.StatusOK)

}