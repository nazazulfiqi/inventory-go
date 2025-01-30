package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name         string   `json:"name" gorm:"type:varchar(200);not null"`
	Description  string   `json:"description" gorm:"type:text"`
	Price        float64  `json:"price" gorm:"type:decimal(10,2);not null"`
	Stock        int      `json:"stock" gorm:"default:0"`
	CategoryID   uint     `json:"category_id" gorm:"not null"`
	Category     Category `json:"-" gorm:"foreignKey:CategoryID"`
	SerialNumber string   `json:"serial_number" gorm:"type:varchar(100);unique"`
}
