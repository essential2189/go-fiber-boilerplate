package model

import "time"

type Purchase struct {
	Uid              uint64    `gorm:"column:uid;primaryKey"`
	BillId           uint64    `gorm:"column:bill_id"`
	PgId             uint16    `gorm:"column:pg_id;default:1"`
	PurchaseMethodId uint16    `gorm:"column:purchase_method_id;default:2"`
	Status           int8      `gorm:"column:status;type:tinyint(2);default:0"`
	TransactionId    string    `gorm:"column:transaction_id;type:varchar(50)"`
	TransactionDate  time.Time `gorm:"column:transaction_date;default:current_timestamp"`
	PgAmount         float32   `gorm:"column:pg_amount;decimal(15,4);default:0"`
	CoinAmount       float32   `gorm:"column:coin_amount;decimal(15,4);default:0"`
	BonusAmount      float32   `gorm:"column:bonus_amount;decimal(15,4);default:0"`
	DiscountAmount   float32   `gorm:"column:discount_amount;decimal(15,4);default:0"`
	Memo             string    `gorm:"column:memo;type:varchar(500)"`
	CreateTs         time.Time `gorm:"column:create_ts;default:current_timestamp"`
}

func (p Purchase) TableName() string {
	return "tb_purchase"
}
