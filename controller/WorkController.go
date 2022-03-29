package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"gatelligance/service"
	"gatelligance/utils"
	Verification "gatelligance/verification"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitWorkController(err *error, db *gorm.DB, router *gin.Engine) {

	router.POST("/frontEnd/checkLinkTransaction", func(c *gin.Context) {
		tid := c.DefaultPostForm("tid", "nil")
		token := c.DefaultPostForm("token", "nil")

		if tid == "nil" || token == "nil" {
			c.String(http.StatusNotAcceptable, fmt.Sprintln("network"))
			return
		}

		success, user := Verification.GetUserFromToken(token, err, db, router)
		if success {
			res := service.CheckLinkTransactionService(tid, db)
			c.JSON(http.StatusOK, utils.CheckLinkTransactionResponse{
				IsSuccess: true,
				ErrorMsg:  "200",
				Progress:  res.Progress,
				Status:    res.Status,
				Output:    res.Output,
				Avatar:    res.Avatar,
				Title:     res.Title,
				Type:      res.Type,
			})
			print(user.ID)
		} else {
			c.JSON(http.StatusOK, utils.CheckLinkTransactionResponse{
				IsSuccess: false,
				ErrorMsg:  "501",
				Progress:  "-1",
				Status:    "-1",
				Output:    "nil",
			})
		}

	})

	router.POST("/frontEnd/fetchList", func(c *gin.Context) {

		token := c.DefaultPostForm("token", "nil")
		page := c.DefaultPostForm("page", "nil")

		if token == "nil" || page == "nil" {
			c.String(http.StatusNotAcceptable, fmt.Sprintln("network"))
			return
		}

		success, user := Verification.GetUserFromToken(token, err, db, router)
		if success {
			pageInt, _ := strconv.Atoi(page)
			var taskList []utils.TaskListRow = service.GetUsersTransactionList(db, user.ID, pageInt)
			c.JSON(200, utils.TransactionListResponse{
				IsSuccess: true,
				ErrorMsg:  "200",
				TaskList:  taskList,
			})
		} else {
			c.JSON(200, utils.TransactionListResponse{
				IsSuccess: false,
				ErrorMsg:  "501",
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
			// Service.CreateLinkTransaction(link, db, user.ID)
			c.JSON(http.StatusOK, utils.WorkSubmitResponse{
				IsSuccess:     true,
				ErrorMsg:      "200",
				TransactionID: service.CreateLinkTransaction(link, db, user.ID),
				// TaskList: Service.CreateLinkTransaction(link),
			})
			// println(user.ID)
		} else {
			c.JSON(http.StatusOK, utils.WorkSubmitResponse{
				IsSuccess:     false,
				ErrorMsg:      "501",
				TransactionID: "",
			})
		}

	})

}
