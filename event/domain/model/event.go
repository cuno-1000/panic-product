package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type Event struct {
	ID                          uint
	Uuid                        string
	Info                        string
	AdminId                     uint
	ApplyRules                  string
	ProductQuantity             uint32          `json:"product_quantity" sql:"type:decimal(20,2);"`
	ProductItemPrice            decimal.Decimal `json:"product_item_price" sql:"type:decimal(20,2);"`
	CreatedAt                   time.Time
	RepaymentReviewUpperLimitAt time.Time
	OverDueMaxTimes             uint
	StartingAt                  time.Time
}

type EventCache struct {
}
