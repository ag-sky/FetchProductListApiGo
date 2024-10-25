package main

import (
	"FetchProductDetail/models"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

var products = []models.Product{
	{ID: 1, Name: "Gatorade",Price: 20.40},
	{ID: 2, Name: "milk",Price: 60.40},
	{ID: 3, Name: "egg",Price: 1.40},
}

func GetProducts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/api/products",GetProducts).Methods("GET")
	http.ListenAndServe(":8000",router)
}