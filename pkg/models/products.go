package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name           string
	URL            string
	LastChecked    string
	UserID         uint
	PriceHistories []PriceHistory
}
