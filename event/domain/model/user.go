package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Tel          string `gorm:"unique_index;not_null"`
	Name         string `gorm:"not_null"`
	IdNumber     string `gorm:"not_null"`
	HashPassword string `gorm:"not_null"`

	CareerStatus string
}
