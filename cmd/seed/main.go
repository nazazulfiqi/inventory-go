package main

import (
	"go-crud/db"
	"go-crud/seeders"
)

func main() {
	db.ConnectDB()

	seeders.SeedCategories()
	seeders.SeedProducts()
}
