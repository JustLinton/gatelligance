package entity

import "github.com/jinzhu/gorm"

type AvatarResourceTable struct {
	ID  int
	Url string
}

func InitAvatarResources(db *gorm.DB) {
	db.AutoMigrate(&AvatarResourceTable{})
	var uu []AvatarResourceTable
	db.Find(&uu, "id=?", 1)
	if len(uu) == 0 {
		//1的意思是count
		tmp := AvatarResourceTable{1, "https://linton-pics.oss-cn-beijing.aliyuncs.com/avatars/dog1.jpg"}
		db.Create(tmp)
	}
}
