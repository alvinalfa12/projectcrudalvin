package handlers

import (
	"encoding/json"
	"go-crud/database"
	"go-crud/models"
	"net/http"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var cat models.Category
	json.NewDecoder(r.Body).Decode(&cat)

	_, err := database.DB.Exec("INSERT INTO categories(name) VALUES(?)", cat.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
