package main

import (
	"apis/product_api"
	"fmt"
	"mux-master"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/product/findall", product_api.FindAll).Methods("GET")
	router.HandleFunc("/api/product/create", product_api.Create).Methods("POST")
	router.HandleFunc("/api/product/update", product_api.Update).Methods("PUT")
	router.HandleFunc("/api/product/delete/{id}", product_api.Delete).Methods("DELETE")
	router.HandleFunc("/api/product/searchid/{id}", product_api.SearchID).Methods("GET")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}

}
