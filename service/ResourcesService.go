package service

import (
	"gatelligance/entity"

	"github.com/jinzhu/gorm"
)

func GetResourceUrl(db *gorm.DB) string {
	var uu []entity.ResourcesTable
	db.Find(&uu, "Label=?", "android-apk")
	if len(uu) != 0 {
		return uu[0].Url
	}
	return "https://a.cupof.beer/"
}
