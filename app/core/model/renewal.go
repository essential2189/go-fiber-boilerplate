package model

import "time"

type Renewal struct {
	Uid           uint64    `gorm:"column:uid;primaryKey"`
	AccountId     uint64    `gorm:"column:account_id"`
	PaymentId     int32     `gorm:"column:payment_Id"`
	ProductId     int32     `gorm:"column:product_id"`
	Type          int8      `gorm:"column:type;type:tinyint(2);default:1"`
	BillId        uint64    `gorm:"column:bill_id"`
	TransactionId string    `gorm:"column:transaction_id;type:varchar(50)"`
	Category      int8      `gorm:"column:category;type:tinyint(2)"`
	Status        int8      `gorm:"column:status;type:tinyint(2);default:0"`
	TargetDate    time.Time `gorm:"column:target_date;type:datetime;default:current_timestamp"`
	RenewalCount  int32     `gorm:"column:renewal_count;default:0"`
	CreateTs      time.Time `gorm:"column:create_ts;default:current_timestamp"`
	UpdateTs      time.Time `gorm:"column:update_ts;on update:current_timestamp"`
}

func (r Renewal) TableName() string {
	return "tb_renewal"
}
