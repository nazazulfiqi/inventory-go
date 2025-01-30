package repository

import (
	"errors"
	"go-crud/db"
	"go-crud/models"
	"time"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := db.DB.Find(&users)
	return users, result.Error
}

func CreateUser(user models.User) error {
	var existingUser models.User

	// Cek apakah email sudah pernah dipakai, termasuk yang sudah soft deleted
	if err := db.DB.Unscoped().Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		// Jika user ditemukan dan sudah dihapus (soft delete), maka restore
		if existingUser.DeletedAt.Valid {
			existingUser.DeletedAt.Time = time.Time{} // Reset deleted_at agar NULL
			existingUser.DeletedAt.Valid = false      // Pastikan NULL

			existingUser.Name = user.Name // Update data jika perlu

			return db.DB.Save(&existingUser).Error // Simpan perubahan
		}

		// Jika email masih aktif, tolak request
		return errors.New("email already exists")
	}

	// Jika email belum ada, buat user baru
	return db.DB.Create(&user).Error
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	result := db.DB.First(&user, id)
	return user, result.Error
}

func UpdateUser(user models.User) error {
	result := db.DB.Save(&user)
	return result.Error
}

func DeleteUser(id uint) error {
	result := db.DB.Delete(&models.User{}, id)
	return result.Error
}

// user_repository.go
func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	result := db.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}
