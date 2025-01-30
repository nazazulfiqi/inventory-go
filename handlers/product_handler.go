package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"go-crud/models"
	"go-crud/repository"

	"github.com/gorilla/mux"
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

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	// Import "github.com/gorilla/mux" dan "strconv" jika belum
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := repository.GetProductByID(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to get product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var updatedProduct models.Product

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Update product
	if err := repository.UpdateProduct(uint(id), updatedProduct); err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product updated"})
}

// Delete Product Handler
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// Delete product
	if err := repository.DeleteProduct(uint(id)); err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted"})
}
