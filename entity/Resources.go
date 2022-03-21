package entity

import "github.com/jinzhu/gorm"

type ResourcesTable struct {
	Label string
	Url   string
}

func InitResources(db *gorm.DB) {
	db.AutoMigrate(&ResourcesTable{})
	var uu []ResourcesTable
	db.Find(&uu, "Label=?", "android-apk")
	if len(uu) == 0 {
		//1的意思是count
		tmp := ResourcesTable{"android-apk", "https://linton-pics.oss-cn-beijing.aliyuncs.com/images/LightShadow-alpha.apk"}
		db.Create(tmp)
	}
}
