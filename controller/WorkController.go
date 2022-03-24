package controller

import (
	"fmt"
	"net/http"

	Service "gatelligance/service"
	Utils "gatelligance/utils"
	Verification "gatelligance/verification"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitWorkController(err *error, db *gorm.DB, router *gin.Engine) {

	router.POST("/frontEnd/uploadLink", func(c *gin.Context) {

		link := c.DefaultPostForm("link", "nil")
		token := c.DefaultPostForm("token", "nil")

		if link == "nil" || token == "nil" {
			c.String(http.StatusNotAcceptable, fmt.Sprintln("network"))
			return
		}

		// claim, stat := Verification.VerifyToken(strToken)
		// if !stat {
		// 	c.String(http.StatusOK, "Login expired.")
		// 	return
		// }
		// c.String(http.StatusOK, "Hello,"+claim.ID)

		success, user := Verification.GetUserFromToken(token, err, db, router)
		if success {
			c.JSON(http.StatusOK, Utils.WorkSubmitResponse{
				IsSuccess: "true",
				ErrorMsg:  "200",
				TaskList:  Service.GetAudioSummary(link, "1"),
			})
			print(user.ID)
		} else {
			c.JSON(http.StatusOK, Utils.WorkSubmitResponse{
				IsSuccess: "false",
				ErrorMsg:  "501",
				TaskList:  "",
			})
		}

	})

}
