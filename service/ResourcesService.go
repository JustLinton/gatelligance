package service

import (
	"gatelligance/entity"

	"github.com/jinzhu/gorm"
)

func GetAvatarResourceUrl(db *gorm.DB, id int) string {
	var uu []entity.AvatarResourceTable
	db.Find(&uu, "id=?", id)
	if len(uu) != 0 {
		return uu[0].Url
	}
	return "https://linton-pics.oss-cn-beijing.aliyuncs.com/avatars/dog1.jpg"
}

func GetAvatarResourceList(db *gorm.DB) []entity.AvatarResourceTable {
	var uu []entity.AvatarResourceTable
	db.Find(&uu)
	return uu
}
