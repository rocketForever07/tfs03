package orderDetailController

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



func GetDetail(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	key := vars["id"]
	id ,_ := strconv.ParseInt(key,10,64)


	orderDetail,err := repo.GetOrderDetailById(id)
	
	fmt.Println(orderDetail)
	if err != nil{
		w.WriteHeader(http.StatusNoContent)
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orderDetail)

}


func Create(w http.ResponseWriter, r *http.Request){

	reqBody,err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Println(err)
	}

	var orderDetail entity.OrderDetail
	//convert json to obj
	json.Unmarshal(reqBody,&orderDetail)
	fmt.Println(orderDetail)

	//import to db
	newOrderDetail,err := repo.CreateOrderDetail(orderDetail)

	if err != nil{
		log.Println(err)	
		w.WriteHeader(http.StatusNoContent)
	}
	//
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	//response this payment obj(json)
	json.NewEncoder(w).Encode(newOrderDetail)

}

func Update(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id,_ := strconv.Atoi(vars["id"])

	requestBody, _ := ioutil.ReadAll(r.Body)
	var orderDetail entity.OrderDetail
	json.Unmarshal(requestBody, &orderDetail)
	orderDetail.ID = id
	log.Println(orderDetail)

	updateOrderDetail,err := repo.UpdateOrderDetail(orderDetail)
	if err != nil{
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updateOrderDetail)

}

func Delete(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	key := vars["id"]

	id ,_ := strconv.ParseInt(key,10,64)

	err := repo.DeleteOrderDetail(id)
	if err!=nil{
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
}