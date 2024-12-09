package entity

import (
	"gorm.io/gorm"
)

// Users 表结构
type Users struct {
	gorm.Model
	Address  string `gorm:"type:varchar(255);not null;unique" json:"address"` // 用户链上地址
	Username string `gorm:"type:varchar(100);unique" json:"username"`         // 用户名
	Password string `gorm:"type:varchar(255);" json:"password"`               // 用户密码（加密后）
	Email    string `gorm:"type:varchar(255);unique" json:"email"`            // 邮箱
	Phone    string `gorm:"type:varchar(50)" json:"phone"`                    // 电话（预留）
	TierID   uint   `gorm:"default:1" json:"tier_id"`                         // 等级 1-5
}

// Assets 表结构
type Assets struct {
	gorm.Model
	UserId     string `gorm:"not null;index" json:"user_id"`                 // 外键，关联 Users 表 一个user 有不同 chain 下的 asset
	ChainID    string `gorm:"type:varchar(100);not null" json:"chain_id"`    // 链ID
	MainPvKey  string `gorm:"type:text;not null" json:"main_pvkey"`          // 主钱包私钥
	MnemonicID string `gorm:"type:varchar(255);not null" json:"mnemonic_id"` // 助记词ID
	Status     int    `gorm:"default:1" json:"status"`                       // 状态字段 1 active，0 inactive
}
