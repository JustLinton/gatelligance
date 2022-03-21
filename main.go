package main

import (
	"gatelligance/controller"
	Entity "gatelligance/entity"
	"log"
	"net/http"

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

	// Service.SendTest()

	router := initAPIs(&err, db)
	router.Use(Cors())
	router.Run(":8082")

}

func initAPIs(err *error, db *gorm.DB) *gin.Engine {
	//start up the gin frame and implement the methods to respond the http requests.
	router := gin.Default()

	controller.InitUsersController(err, db, router)

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "这里是凝智成林业务后端👋")
	})

	return router
}

func connectDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "backend:As123456$@(rm-wz9637tu57d99e8665o.mysql.rds.aliyuncs.com:3306)/gatelligance?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("open database error:%v", err)
		return nil, err
	}
	initEntities(db)
	return db, nil
}

func initEntities(db *gorm.DB) {
	Entity.InitUsers(db)
	Entity.InitResources(db)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
