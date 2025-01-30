package repository

import (
	"go-crud/db"
	"go-crud/models"
)

func CreateProduct(product models.Product) error {
	return db.DB.Create(&product).Error
}

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	result := db.DB.Preload("Category").Find(&products)
	return products, result.Error
}

func GetProductByID(id uint) (models.Product, error) {
	var product models.Product
	result := db.DB.Preload("Category").First(&product, id)
	return product, result.Error
}

func UpdateProduct(id uint, updatedProduct models.Product) error {
	var product models.Product

	// Cek apakah produk ada
	if err := db.DB.First(&product, id).Error; err != nil {
		return err
	}

	// Update produk
	product.Name = updatedProduct.Name
	product.Description = updatedProduct.Description
	product.Price = updatedProduct.Price
	product.CategoryID = updatedProduct.CategoryID // Update kategori jika perlu

	return db.DB.Save(&product).Error
}

// Delete Product
func DeleteProduct(id uint) error {
	return db.DB.Delete(&models.Product{}, id).Error
}
