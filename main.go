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
	db.DB.AutoMigrate(&models.User{})

	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Go CRUD API")
	}).Methods("GET")
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	// Protected routes
	authRouter := r.PathPrefix("/").Subrouter()
	authRouter.Use(handlers.AuthMiddleware)

	authRouter.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	authRouter.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	authRouter.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	authRouter.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	authRouter.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
