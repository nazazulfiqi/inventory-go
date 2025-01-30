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

func GetCategoryByID(id uint) (models.Category, error) {
	var category models.Category
	result := db.DB.First(&category, id)
	return category, result.Error
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

// Update Category
func UpdateCategory(id uint, updatedCategory models.Category) error {
	var category models.Category

	// Cek apakah kategori ada
	if err := db.DB.First(&category, id).Error; err != nil {
		return err
	}

	// Update kategori
	category.Name = updatedCategory.Name
	category.Description = updatedCategory.Description

	return db.DB.Save(&category).Error
}

// Delete Category
func DeleteCategory(id uint) error {
	return db.DB.Delete(&models.Category{}, id).Error
}
