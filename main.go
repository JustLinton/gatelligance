package main

import (
	"gatelligance/controller"
	Entity "gatelligance/entity"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := connectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := initAPIs(&err, db)
	//设置跨域
	router.Use(cors.Default())
	router.Run(":8082")
}

func initAPIs(err *error, db *gorm.DB) *gin.Engine {
	//start up the gin frame and implement the methods to respond the http requests.
	router := gin.Default()

	controller.InitUsersApi(db, router)

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "欢迎访问轻影👋")
	})

	return router
}

func connectDatabase() (*gorm.DB, error) {
	//connect the database.
	//db, err := gorm.Open("mysql", "backend:123456@(127.0.0.1:3306)/lightshadow?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", "backend:As123456$@(rm-wz9637tu57d99e8665o.mysql.rds.aliyuncs.com:3306)/lightshadow?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("open database error:%v", err)
		return nil, err
	}
	initEntities(db)
	return db, nil
}

func initEntities(db *gorm.DB) {
	Entity.InitUser(db)
	Entity.InitResources(db)
}
