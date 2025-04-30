package main

import (
	"go-crud/database"
	"go-crud/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.Connect()

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	r.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")
	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
