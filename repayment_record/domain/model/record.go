package model

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"time"
)

type Record struct {
	gorm.Model
	DueRepaymentTimestamp    time.Time `gorm:"not_null"`
	ActualRepaymentTimestamp time.Time `gorm:"not_null"`
	DateDifference           int       `gorm:"not_null"`
	IsOnTime                 bool
	OwedAmount               decimal.Decimal `json:"amount" sql:"type:decimal(20,2);"`
	UserId                   uint
	User                     User
}
