package model

import "time"

type WalletProduct struct {
	Uid           uint64    `gorm:"column:uid;primaryKey"`
	WalletId      uint64    `gorm:"column:wallet_id"`
	ProductId     int32     `gorm:"column:product_id"`
	BillId        uint64    `gorm:"column:bill_id"`
	Type          int8      `gorm:"column:type;type:tinyint(2);default:1"`
	Priority      int8      `gorm:"column:priority;type:tinyint(2);default:1"`
	Renew         int8      `gorm:"column:renew;type:tinyint(2);default:1"`
	StartTs       time.Time `gorm:"column:start_ts;type:datetime;default:current_timestamp"`
	EndTs         time.Time `gorm:"column:end_ts;type:datetime;default:current_timestamp"`
	RemainCount   uint64    `gorm:"column:remain_count;default:0"`
	DownloadCount uint16    `gorm:"column:download_count;default:0"`
	Status        int8      `gorm:"column:status;type:tinyint(2);default:0"`
	CreateTs      time.Time `gorm:"column:create_ts;default:current_timestamp"`
	UpdateTs      time.Time `gorm:"column:update_ts;on update:current_timestamp"`
}

func (w WalletProduct) TableName() string {
	return "tb_wallet_product"
}
