package handlers

import (
	"encoding/json"
	"net/http"

	"go-crud/models"
	"go-crud/repository"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := repository.CreateCategory(category); err != nil {
		http.Error(w, "Failed to create category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := repository.GetAllCategories()
	if err != nil {
		http.Error(w, "Failed to get categories", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)
}
