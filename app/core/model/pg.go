package model

import "time"

type Pg struct {
	Uid      uint16    `gorm:"column:uid;primaryKey"`
	Name     string    `gorm:"column:name;type:varchar(255)"`
	CreateTs time.Time `gorm:"column:create_ts;default:current_timestamp"`
}

func (p Pg) TableName() string {
	return "tb_pg"
}
