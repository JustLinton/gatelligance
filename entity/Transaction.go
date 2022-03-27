package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Transaction struct {
	ID        string `gorm:"primary_key"`
	Server    int
	Owner     string
	CreatedAt time.Time
}

func InitTransaction(db *gorm.DB) {
	db.AutoMigrate(&Transaction{})
}
