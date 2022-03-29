package entity

import (
	"github.com/jinzhu/gorm"
)

type EmailActiCode struct {
	Uuid string `gorm:"primary_key"`
	Code string
}

func InitEmailActiCode(db *gorm.DB) {
	db.AutoMigrate(&EmailActiCode{})
}
