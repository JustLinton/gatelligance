package entity

import "github.com/jinzhu/gorm"

// Transaction 根据链接提取要素的算法事务信息
type LinkTransaction struct {
	ID        string `gorm:"primary_key"`
	VideoLink string
	Progress  string //单位：%
	Status    string //0-doing,1-ok,-1-error
	Output    string
}

func InitLinkTransactionEntity(db *gorm.DB) {
	db.AutoMigrate(&LinkTransaction{})
}
