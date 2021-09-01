package categoryController

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

	categories,err := repo.GetAllCategory()

	if err != nil{
		w.WriteHeader(http.StatusNoContent)
		log.Println(err)
	}else{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	
		json.NewEncoder(w).Encode(categories)
	}

}


func GetDetail(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	key := vars["id"]
	id ,_ := strconv.ParseInt(key,10,64)


	category,err := repo.GetCategoryById(id)
	
	fmt.Println(category)
	if err != nil{
		w.WriteHeader(http.StatusNoContent)
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)

}


func Create(w http.ResponseWriter, r *http.Request){

	reqBody,err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Println(err)
	}

	var category entity.Category
	//convert json to obj
	json.Unmarshal(reqBody,&category)
	fmt.Println(category)

	//import to db
	newCategory,err := repo.CreateCategory(category)

	if err != nil{
		log.Println(err)	
		w.WriteHeader(http.StatusNoContent)
	}
	//
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	//response this payment obj(json)
	json.NewEncoder(w).Encode(newCategory)

}

func Update(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id,_ := strconv.Atoi(vars["id"])

	requestBody, _ := ioutil.ReadAll(r.Body)
	var category entity.Category
	json.Unmarshal(requestBody, &category)
	category.ID = id
	log.Println(category)

	updateCategory,err := repo.UpdateCategory(category)
	if err != nil{
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updateCategory)

}

func Delete(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	key := vars["id"]

	id ,_ := strconv.ParseInt(key,10,64)

	err := repo.DeleteCategory(id)
	if err!=nil{
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
}