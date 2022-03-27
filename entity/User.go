package entity

import "github.com/jinzhu/gorm"

// User UserInfo 用户信息
type User struct {
	ID       string `gorm:"primary_key"`
	Phone    string
	NickName string
	Email    string
	PassSHA  string
	Gender   string
	Avatar   string
}

func InitUsers(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
