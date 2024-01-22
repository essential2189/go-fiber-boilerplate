package model

import "time"

type Bill struct {
	Uid         uint64    `gorm:"column:uid;primaryKey"`
	AccountId   uint64    `gorm:"column:account_id"`
	PaymentId   int32     `gorm:"column:payment_id"`
	CurrencyId  uint16    `gorm:"column:currency_id"`
	Type        int8      `gorm:"column:type;type:tinyint(2);default:1"`
	DeviceAgent string    `gorm:"column:device_agent;type:text"`
	TotalAmount float32   `gorm:"column:total_amount;default:0"`
	CreateTs    time.Time `gorm:"column:create_ts;default:current_timestamp"`
}

func (b Bill) TableName() string {
	return "tb_bill"
}
