// repository/category_repository.go
package repository

import (
	"go-crud/db"
	"go-crud/models"
)

func CreateCategory(category models.Category) error {
	return db.DB.Create(&category).Error
}

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	result := db.DB.Find(&categories)
	return categories, result.Error
}

func CategoryExists(id uint) (bool, error) {
	var category models.Category
	result := db.DB.First(&category, id)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}
