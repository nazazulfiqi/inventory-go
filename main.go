package main

import (
	"fmt"
	"go-crud/db"
	"go-crud/handlers"
	"go-crud/models"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.ConnectDB()
	db.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Category{})

	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Go CRUD API")
	}).Methods("GET")
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	// Protected routes (untuk semua role terautentikasi)
	authRouter := r.PathPrefix("/").Subrouter()
	authRouter.Use(handlers.AuthMiddleware)
	authRouter.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	authRouter.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")

	// Routes khusus admin dan supervisor
	adminSupervisorRouter := authRouter.PathPrefix("/").Subrouter()
	adminSupervisorRouter.Use(handlers.SupervisorMiddleware) // SupervisorMiddleware allows both admin & supervisor

	// User management routes
	adminSupervisorRouter.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	adminSupervisorRouter.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	adminSupervisorRouter.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	// Admin routes
	adminRouter := r.PathPrefix("/admin").Subrouter()
	adminRouter.Use(handlers.AuthMiddleware)
	adminRouter.Use(handlers.AdminMiddleware)

	// Product routes
	adminRouter.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	// adminRouter.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	// adminRouter.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")

	// Category routes
	adminRouter.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")
	// adminRouter.HandleFunc("/categories/{id}", handlers.UpdateCategory).Methods("PUT")
	// adminRouter.HandleFunc("/categories/{id}", handlers.DeleteCategory).Methods("DELETE")

	// Public read routes
	r.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	// r.HandleFunc("/products/{id}", handlers.GetProduct).Methods("GET")
	r.HandleFunc("/categories", handlers.GetCategories).Methods("GET")

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
