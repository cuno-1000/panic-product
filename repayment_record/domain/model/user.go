package model

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type User struct {
	gorm.Model
	Tel            string          `gorm:"unique_index;not_null"`
	Name           string          `gorm:"not_null"`
	IdNumber       string          `gorm:"not_null"`
	HashPassword   string          `gorm:"not_null"`
	AccountBalance decimal.Decimal `json:"product_item_price" sql:"type:decimal(20,2);"`

	CareerStatus string
}
