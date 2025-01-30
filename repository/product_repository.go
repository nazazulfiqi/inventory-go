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

func UpdateProduct(product models.Product) error {
	return db.DB.Save(&product).Error
}

func DeleteProduct(id uint) error {
	return db.DB.Delete(&models.Product{}, id).Error
}
