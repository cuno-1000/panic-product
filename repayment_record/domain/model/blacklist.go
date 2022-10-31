package model

import "github.com/jinzhu/gorm"

type BlackList struct {
	gorm.Model
	Reason string
	UserId uint
	User   User
}
