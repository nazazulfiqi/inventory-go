package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);unique;not null"`
	Description string `gorm:"type:text"`
	Products    []Product
}
