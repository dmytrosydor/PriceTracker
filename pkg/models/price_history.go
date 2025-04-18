package models

import "gorm.io/gorm"

type PriceHistory struct {
	gorm.Model
	ProductID uint
	Price     float64
	CheckedAt string
}
