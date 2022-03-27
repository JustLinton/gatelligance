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

	router.POST("/frontEnd/checkLinktransaction", func(c *gin.Context) {
		tid := c.DefaultPostForm("tid", "nil")
		token := c.DefaultPostForm("token", "nil")

		if tid == "nil" || token == "nil" {
			c.String(http.StatusNotAcceptable, fmt.Sprintln("network"))
			return
		}

		success, user := Verification.GetUserFromToken(token, err, db, router)
		if success {
			prog, stat, output := Service.CheckLinkTransactionService(tid, db)
			c.JSON(http.StatusOK, Utils.CheckLinkTransactionResponse{
				IsSuccess: "true",
				ErrorMsg:  "200",
				Progress:  prog,
				Status:    stat,
				Output:    output,
			})
			print(user.ID)
		} else {
			c.JSON(http.StatusOK, Utils.CheckLinkTransactionResponse{
				IsSuccess: "false",
				ErrorMsg:  "501",
				Progress:  "-1",
				Status:    "-1",
				Output:    "nil",
			})
		}

	})

	router.POST("/frontEnd/uploadLink", func(c *gin.Context) {

		link := c.DefaultPostForm("link", "nil")
		token := c.DefaultPostForm("token", "nil")

		if link == "nil" || token == "nil" {
			c.String(http.StatusNotAcceptable, fmt.Sprintln("network"))
			return
		}

		success, user := Verification.GetUserFromToken(token, err, db, router)
		if success {
			c.JSON(http.StatusOK, Utils.WorkSubmitResponse{
				IsSuccess: "true",
				ErrorMsg:  "200",
				TaskList:  Service.CreateLinkTransaction(link, db, user.ID),
				// TaskList: Service.CreateLinkTransaction(link),
			})
			// println(user.ID)
		} else {
			c.JSON(http.StatusOK, Utils.WorkSubmitResponse{
				IsSuccess: "false",
				ErrorMsg:  "501",
				TaskList:  "",
			})
		}

	})

}
