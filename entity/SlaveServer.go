package entity

import (
	"github.com/jinzhu/gorm"
)

type SlaveServer struct {
	ServerID int `gorm:"primary_key"`
	Type     string
	Address  string
	Usable   int
}

func InitSlaveServer(db *gorm.DB) {
	db.AutoMigrate(&SlaveServer{})
}
