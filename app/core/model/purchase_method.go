package model

import "time"

type PurchaseMethod struct {
	Uid      uint16    `gorm:"column:uid;primaryKey"`
	Name     string    `gorm:"column:name;type:varchar(255)"`
	CreateTs time.Time `gorm:"column:create_ts;default:current_timestamp"`
}

func (p PurchaseMethod) TableName() string {
	return "tb_purchase_method"
}
