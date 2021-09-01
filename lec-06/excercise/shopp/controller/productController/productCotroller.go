package productController

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

	products,err := repo.GetAllProduct()

	if err != nil{
		w.WriteHeader(http.StatusNoContent)
		log.Println(err)
	}else{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	
		json.NewEncoder(w).Encode(products)
	}

}


func GetDetail(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	key := vars["id"]
	id ,_ := strconv.ParseInt(key,10,64)


	product,err := repo.GetProductById(id)
	
	fmt.Println(product)
	if err != nil{
		w.WriteHeader(http.StatusNoContent)
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)

}


func Create(w http.ResponseWriter, r *http.Request){

	reqBody,err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Println(err)
	}

	var product entity.Product
	//convert json to obj
	json.Unmarshal(reqBody,&product)
	fmt.Println(product)

	//import to db
	newProduct,err := repo.CreateProduct(product)

	if err != nil{
		log.Println(err)	
		w.WriteHeader(http.StatusNoContent)
	}
	//
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	//response this payment obj(json)
	json.NewEncoder(w).Encode(newProduct)

}

func Update(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id,_ := strconv.Atoi(vars["id"])

	requestBody, _ := ioutil.ReadAll(r.Body)
	var product entity.Product
	json.Unmarshal(requestBody, &product)
	product.ID = id
	log.Println(product)

	updateProduct,err := repo.UpdateProduct(product)
	if err != nil{
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updateProduct)

}

func Delete(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	key := vars["id"]

	id ,_ := strconv.ParseInt(key,10,64)

	err := repo.DeleteProduct(id)
	if err!=nil{
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
}