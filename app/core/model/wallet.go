package model

import "time"

type Wallet struct {
	Uid       uint64    `gorm:"column:uid;primaryKey"`
	AccountId uint64    `gorm:"column:account_id"`
	Name      string    `gorm:"column:name;type:varchar(200)"`
	IsUse     int8      `gorm:"column:is_use;type:tinyint(2);default:0"`
	CreateTs  time.Time `gorm:"column:create_ts;default:current_timestamp"`
	UpdateTs  time.Time `gorm:"column:update_ts;on update:current_timestamp"`
}

func (w Wallet) TableName() string {
	return "tb_wallet"
}
