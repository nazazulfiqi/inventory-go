package seeders

import (
	"fmt"
	"go-crud/db"
	"go-crud/models"
	"math/rand"
	"time"
)

func SeedCategories() {
	categories := []models.Category{
		{Name: "Laptop", Description: "Berbagai jenis laptop dan notebook"},
		{Name: "Smartphone", Description: "Ponsel pintar dan aksesoris"},
		{Name: "Komponen PC", Description: "Processor, VGA, Motherboard dll"},
		{Name: "Aksesoris", Description: "Kabel, Charger, Case dll"},
	}

	for _, category := range categories {
		db.DB.FirstOrCreate(&category, models.Category{Name: category.Name})
	}
	fmt.Println("Seeded categories successfully!")
}

func SeedProducts() {
	var categories []models.Category
	db.DB.Find(&categories)

	products := []models.Product{
		{
			Name:         "Asus ROG Zephyrus G14",
			Description:  "Laptop gaming AMD Ryzen 9 6900HS, RTX 3060",
			Price:        19999000,
			Stock:        15,
			CategoryID:   findCategoryID(categories, "Laptop"),
			SerialNumber: "ASUS-ROG-G14-2023",
		},
		{
			Name:         "iPhone 15 Pro Max",
			Description:  "Smartphone flagship Apple dengan kamera 48MP",
			Price:        24999000,
			Stock:        30,
			CategoryID:   findCategoryID(categories, "Smartphone"),
			SerialNumber: "APPLE-IP15-PROMAX",
		},
		{
			Name:         "Logitech MX Master 3S",
			Description:  "Mouse wireless premium untuk produktivitas",
			Price:        1499000,
			Stock:        50,
			CategoryID:   findCategoryID(categories, "Aksesoris"),
			SerialNumber: "LOG-MX-M3S-2023",
		},
		// Tambahkan lebih banyak produk sesuai kebutuhan
	}

	// Generate 20 random products
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		products = append(products, models.Product{
			Name:         fmt.Sprintf("Product Demo %d", i+1),
			Description:  fmt.Sprintf("Deskripsi produk demo %d", i+1),
			Price:        float64(rand.Intn(5000000-500000)+500000) + 0.99,
			Stock:        rand.Intn(100),
			CategoryID:   categories[rand.Intn(len(categories))].ID,
			SerialNumber: fmt.Sprintf("SN-%d-%s", time.Now().Unix(), randomString(6)),
		})
	}

	for _, product := range products {
		db.DB.FirstOrCreate(&product, models.Product{SerialNumber: product.SerialNumber})
	}
	fmt.Println("Seeded products successfully!")
}

func findCategoryID(categories []models.Category, name string) uint {
	for _, c := range categories {
		if c.Name == name {
			return c.ID
		}
	}
	return 1
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
