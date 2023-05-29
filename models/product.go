package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	// Id          int64  `gorm:"primaryKey"json:"id"`
	ProductName string `gorm:"type:varchar(360)"json:"product_name"`
	Description string `gorm:"type:text"json:"description"`
}
