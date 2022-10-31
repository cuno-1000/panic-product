package model

//type Order struct {
//	gorm.Model
//	PaymentAmount decimal.Decimal `json:"payment_amount" sql:"type:decimal(20,2);"`
//	EventId       uint
//	UserId        uint
//	ApplyStatus   bool
//}

type EventApplyRecord struct {
	EventId   uint64 `json:"event_id"`
	UserId    uint64 `json:"user_id"`
	Situation int64  `json:"situation"`
}
