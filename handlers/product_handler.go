package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go-crud/models"
	"go-crud/repository"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validasi category_id
	if product.CategoryID == 0 {
		http.Error(w, "Category ID is required", http.StatusBadRequest)
		return
	}

	// Optional: Tambahkan validasi apakah kategori exists
	if exists, err := repository.CategoryExists(product.CategoryID); err != nil {
		http.Error(w, "Error checking category: "+err.Error(), http.StatusInternalServerError)
		return
	} else if !exists {
		http.Error(w, fmt.Sprintf("Category with ID %d does not exist", product.CategoryID), http.StatusBadRequest)
		return
	}

	if err := repository.CreateProduct(product); err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			http.Error(w, fmt.Sprintf("Invalid category ID: %d", product.CategoryID), http.StatusBadRequest)
			return
		}
		http.Error(w, "Failed to create product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := repository.GetAllProducts()
	if err != nil {
		http.Error(w, "Failed to get products", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}
