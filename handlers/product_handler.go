package handlers

import (
	"encoding/json"
	"go-crud/database"
	"go-crud/models"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	json.NewDecoder(r.Body).Decode(&p)

	_, err := database.DB.Exec("INSERT INTO products(name, price, category_id) VALUES(?, ?, ?)", p.Name, p.Price, p.CategoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
        SELECT p.id, p.name, p.price, c.id, c.name
        FROM products p
        JOIN categories c ON p.category_id = c.id`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		var c models.Category
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &c.ID, &c.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		p.Category = &c
		products = append(products, p)
	}

	json.NewEncoder(w).Encode(products)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.Product
	json.NewDecoder(r.Body).Decode(&p)

	_, err := database.DB.Exec("UPDATE products SET name=?, price=?, category_id=? WHERE id=?", p.Name, p.Price, p.CategoryID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := database.DB.Exec("DELETE FROM products WHERE id=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
