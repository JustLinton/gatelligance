package service

import (
	"fmt"
	"gatelligance/entity"

	"github.com/jinzhu/gorm"
)

func AddDownloadCount(db *gorm.DB) {

	var uu []entity.DownloadTable
	db.Find(&uu, "field1=?", 1)
	if len(uu) != 0 {
		db.Delete(uu[0])
		uu[0].Field2 = uu[0].Field2 + 1
		db.Create(uu[0])
		fmt.Printf("yes")
	}

}

func GetDownloadCount(db *gorm.DB) int {
	var uu []entity.DownloadTable
	db.Find(&uu, "field1=?", 1)
	if len(uu) != 0 {
		return uu[0].Field2
	}
	return -1
}

func SetDownloadCount(db *gorm.DB, count int) {
	var uu []entity.DownloadTable
	db.Find(&uu, "field1=?", 1)
	if len(uu) != 0 {
		db.Delete(uu[0])
		uu[0].Field2 = count
		db.Create(uu[0])
	}
}
