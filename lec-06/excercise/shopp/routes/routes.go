package routes

import(
	categoryController "shopp/controller/categoryController"
	productController "shopp/controller/productController"
	orderController "shopp/controller/orderController"
	orderDetailController "shopp/controller/orderDetailController"
	cartController "shopp/controller/cartController"
	cartDetailController "shopp/controller/cartDetailController"
	paymentController "shopp/controller/paymentController"

	"github.com/gorilla/mux"
	"net/http"
	"log"
	"fmt"

)

func Init(){


	//main route
	router := mux.NewRouter().StrictSlash(true)
	//payment route
	paymentRoute := router.PathPrefix("/api/payment").Subrouter()

	//product route
	productRoute := router.PathPrefix("/api/product").Subrouter()

	//category route
	categoryRoute := router.PathPrefix("/api/category").Subrouter()

	//order route
	orderRoute := router.PathPrefix("/api/order").Subrouter()

	//order detail route
	orderDetailRoute := router.PathPrefix("/api/orderDetail").Subrouter()

	//cart route
	cartRoute := router.PathPrefix("/api/cart").Subrouter()

	//cart detail route
	cartDetailRoute := router.PathPrefix("/api/cartDetail").Subrouter()


	//handler for main route
	router.HandleFunc("/home",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w,"home page")
	})

	//handle for category route
	categoryRoute.HandleFunc("",categoryController.GetAll).Methods("GET")
	categoryRoute.HandleFunc("/{id}",categoryController.GetDetail).Methods("GET")
	categoryRoute.HandleFunc("",categoryController.Create).Methods("POST")
	categoryRoute.HandleFunc("/{id}",categoryController.Update).Methods("PUT")
	categoryRoute.HandleFunc("/{id}",categoryController.Delete).Methods("DELETE")

	//handle for product route
	productRoute.HandleFunc("",productController.GetAll).Methods("GET")
	productRoute.HandleFunc("/{id}",productController.GetDetail).Methods("GET")
	productRoute.HandleFunc("",productController.Create).Methods("POST")
	productRoute.HandleFunc("/{id}",productController.Update).Methods("PUT")
	productRoute.HandleFunc("/{id}",productController.Delete).Methods("DELETE")

	//handle for payment route
	paymentRoute.HandleFunc("",paymentController.GetAll).Methods("GET")
	paymentRoute.HandleFunc("",paymentController.Create).Methods("POST")


	//handle for order route
	orderRoute.HandleFunc("",orderController.GetAll).Methods("GET")
	orderRoute.HandleFunc("/{id}",orderController.GetDetail).Methods("GET")
	orderRoute.HandleFunc("",orderController.Create).Methods("POST")
	orderRoute.HandleFunc("/{id}",orderController.Update).Methods("PUT")
	orderRoute.HandleFunc("/{id}",orderController.Delete).Methods("DELETE")

	//handle for orderdetail route
	orderDetailRoute.HandleFunc("/{id}",orderDetailController.GetDetail).Methods("GET")
	orderDetailRoute.HandleFunc("",orderDetailController.Create).Methods("POST")
	orderDetailRoute.HandleFunc("/{id}",orderDetailController.Update).Methods("PUT")
	orderDetailRoute.HandleFunc("/{id}",orderDetailController.Delete).Methods("DELETE")

	//handle for cart route
	cartRoute.HandleFunc("",cartController.GetAll).Methods("GET")
	cartRoute.HandleFunc("/{id}",cartController.GetDetail).Methods("GET")
	cartRoute.HandleFunc("",cartController.Create).Methods("POST")
	cartRoute.HandleFunc("/{id}",cartController.Update).Methods("PUT")
	cartRoute.HandleFunc("/{id}",cartController.Delete).Methods("DELETE")

	//handle for cartdetail route
	cartDetailRoute.HandleFunc("/{id}",cartDetailController.GetDetail).Methods("GET")
	cartDetailRoute.HandleFunc("",cartDetailController.Create).Methods("POST")
	cartDetailRoute.HandleFunc("/{id}",cartDetailController.Update).Methods("PUT")
	cartDetailRoute.HandleFunc("/{id}",cartDetailController.Delete).Methods("DELETE")


	log.Println("listen server on port 10000")
	// log.Fatal(http.ListenAndServe(":10000",router))
	http.ListenAndServe(":10000",router)

}

