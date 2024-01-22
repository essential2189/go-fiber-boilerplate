package model

type GetWalletListModel struct {
	// tb_wallet
	IsUsed int `gorm:"column:w.is_use"`

	// tb_wallet_product
	WalletType    int    `gorm:"column:tb_wallet_product.type"`
	Priority      int    `gorm:"column:tb_wallet_product.priority"`
	Renew         int    `gorm:"column:tb_wallet_product.renew"`
	StartTs       string `gorm:"column:tb_wallet_product.start_ts"`
	EndTs         string `gorm:"column:tb_wallet_product.end_ts"`
	DownloadCount int    `gorm:"column:tb_wallet_product.download_count"`
	WalletStatus  int    `gorm:"column:tb_wallet_product.status"`

	// tb_bill
	CurrencyType int     `gorm:"column:b.currency_type"`
	BillType     int     `gorm:"column:b.type"`
	BillStatus   int     `gorm:"column:b.status"`
	DeviceAgent  string  `gorm:"column:b.device_agent"`
	TotalAmount  float64 `gorm:"column:b.total_amount"`
}
