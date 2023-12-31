package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeypc/go-restapi-mux/controllers/productcontroller"
	"github.com/jeypc/go-restapi-mux/models"
)

func main() {
	models.ConnectDatabase()

	// Routing
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/products", productcontroller.Index).Methods("GET")
	r.HandleFunc("/product/{id}", productcontroller.Show).Methods("GET")
	r.HandleFunc("/product", productcontroller.Create).Methods("POST")
	r.HandleFunc("/product/{id}", productcontroller.Update).Methods("PUT")
	r.HandleFunc("/product", productcontroller.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
} 