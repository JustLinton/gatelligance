package entity

import "github.com/jinzhu/gorm"

type DownloadTable struct {
	Field1 int
	Field2 int
}

func InitUser(db *gorm.DB) {
	db.AutoMigrate(&DownloadTable{})
	var uu []DownloadTable
	db.Find(&uu, "field1=?", 1)
	if len(uu) == 0 {
		//1的意思是count
		tmp := DownloadTable{1, 0}
		db.Create(tmp)
	}
}
