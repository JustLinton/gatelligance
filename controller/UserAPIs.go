package controller

import (
	"gatelligance/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Result struct {
	Count int
}

func InitUsersApi(db *gorm.DB, router *gin.Engine) {
	router.GET("/download", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		service.AddDownloadCount(db)
		c.Redirect(302, service.GetResourceUrl(db))
	})

	router.GET("/download-count", func(c *gin.Context) {
		theCount := service.GetDownloadCount(db)
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, Result{
			Count: theCount,
		})
	})

	router.POST("/download-count", func(c *gin.Context) {
		countStr := c.DefaultPostForm("Count", "0")
		count, err := strconv.Atoi(countStr)
		if err != nil {
			c.String(http.StatusBadRequest, "error count.")
		}
		c.Header("Access-Control-Allow-Origin", "*")
		service.SetDownloadCount(db, count)
		c.String(http.StatusOK, "ok")
	})
}
