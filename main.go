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
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Go CRUD API")
	}).Methods("GET")

	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
