package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"gatelligance/entity"
	"gatelligance/service"
	"gatelligance/utils"
	Verification "gatelligance/verification"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitUsersController(err *error, db *gorm.DB, router *gin.Engine) {

	router.POST("/frontEnd/login", func(c *gin.Context) {

		password := c.DefaultPostForm("password", "nil")
		email := c.DefaultPostForm("email", "nil")

		if password == "nil" || email == "nil" {
			// c.String(http.StatusNotAcceptable, fmt.Sprintln("network"))
			c.JSON(http.StatusOK, utils.LoginResponse{
				Token:     "-1",
				IsSuccess: false,
				ErrorMsg:  "100",
			})
			return
		}

		service.HandleUserLogin(password, email, db, c)
	})

	router.POST("/frontEnd/register", func(c *gin.Context) {

		email := c.DefaultPostForm("email", "nil")
		password := c.DefaultPostForm("password", "nil")
		nickName := c.DefaultPostForm("nickName", "nil")

		if email == "nil" || password == "nil" || nickName == "nil" {
			// c.String(http.StatusNotAcceptable, fmt.Sprintln("network"))
			c.JSON(http.StatusOK, utils.RegisterResponse{
				IsSuccess: false,
				ErrorMsg:  "100",
			})
		}

		service.HandleUserRegister(password, email, nickName, db, err, c)
	})

	// router.GET("/frontEnd/refreshToken", Verification.RefreshTokenHandler)

	router.POST("/frontEnd/refreshToken", func(c *gin.Context) {

		token := c.DefaultPostForm("token", "nil")

		if token == "nil" {
			c.String(http.StatusNotAcceptable, fmt.Sprintln("network"))
			return
		}

		success, _ := Verification.GetUserFromToken(token, err, db, router)

		if success {
			newToken := Verification.RefreshToken(token)
			if newToken == "-1" {
				c.JSON(200, utils.RefreshTokenResponse{
					IsSuccess: false,
					ErrorMsg:  "501",
				})
			} else {
				c.JSON(200, utils.RefreshTokenResponse{
					IsSuccess: true,
					ErrorMsg:  "200",
					Token:     newToken,
				})
			}

		} else {
			c.JSON(200, utils.StandardResponse{
				IsSuccess: false,
				ErrorMsg:  "501",
			})
		}

	})

	router.POST("/frontEnd/setUserInfo", func(c *gin.Context) {

		token := c.DefaultPostForm("token", "nil")
		name := c.DefaultPostForm("name", "nil")
		avatarID := c.DefaultPostForm("avatar", "nil")
		gender := c.DefaultPostForm("gender", "nil")
		email := c.DefaultPostForm("email", "nil")

		if token == "nil" {
			c.String(http.StatusNotAcceptable, fmt.Sprintln("network"))
			return
		}

		success, user := Verification.GetUserFromToken(token, err, db, router)
		avatarInt, _ := strconv.Atoi(avatarID)
		if success {
			db.Delete(user)
			user.NickName = name
			user.Gender = gender
			user.Avatar = avatarInt
			if email != user.Email {
				user.Activated = 0
				user.Email = email
			}
			db.Create(user)
			c.JSON(http.StatusOK, utils.SetUserInfoResponse{
				IsSuccess: true,
				ErrorMsg:  "200",
			})
		} else {
			c.JSON(http.StatusOK, utils.SetUserInfoResponse{
				IsSuccess: false,
				ErrorMsg:  "501",
			})
		}

	})

	router.POST("/frontEnd/fetchUserInfo", func(c *gin.Context) {

		token := c.DefaultPostForm("token", "nil")

		if token == "nil" {
			c.String(http.StatusNotAcceptable, fmt.Sprintln("network"))
			return
		}

		success, user := Verification.GetUserFromToken(token, err, db, router)

		var activated = false
		if user.Activated == 1 {
			activated = true
		}

		if success {
			c.JSON(http.StatusOK, utils.FetchUserInfoResponse{
				IsSuccess: true,
				ErrorMsg:  "200",
				Email:     user.Email,
				NickName:  user.NickName,
				Avatar:    service.GetAvatarResourceUrl(db, user.Avatar),
				Gender:    user.Gender,
				Activated: activated,
			})
		} else {
			c.JSON(http.StatusOK, utils.FetchUserInfoResponse{
				IsSuccess: false,
				ErrorMsg:  "501",
			})
		}

	})

	router.POST("/frontEnd/fetchAvatarList", func(c *gin.Context) {

		token := c.DefaultPostForm("token", "nil")

		if token == "nil" {
			c.String(http.StatusNotAcceptable, fmt.Sprintln("network"))
			return
		}

		success, _ := Verification.GetUserFromToken(token, err, db, router)
		if success {
			var avaList []entity.AvatarResourceTable = service.GetAvatarResourceList(db)
			c.JSON(200, utils.AvatarListResponse{
				IsSuccess: true,
				ErrorMsg:  "200",
				Data:      avaList,
			})
		} else {
			c.JSON(200, utils.AvatarListResponse{
				IsSuccess: false,
				ErrorMsg:  "501",
			})
		}

	})

	router.POST("/frontEnd/activateEmail", func(c *gin.Context) {

		token := c.DefaultPostForm("token", "nil")
		code := c.DefaultPostForm("code", "nil")

		if token == "nil" || code == "nil" {
			c.String(http.StatusNotAcceptable, fmt.Sprintln("network"))
			return
		}

		success, user := Verification.GetUserFromToken(token, err, db, router)
		if success {
			if service.ActivateEmail(user.ID, code, db) {
				c.JSON(200, utils.StandardResponse{
					IsSuccess: true,
					ErrorMsg:  "200",
				})
			} else {
				c.JSON(200, utils.StandardResponse{
					IsSuccess: false,
					ErrorMsg:  "0",
				})
			}

		} else {
			c.JSON(200, utils.StandardResponse{
				IsSuccess: false,
				ErrorMsg:  "501",
			})
		}

	})

	router.POST("/frontEnd/sendActivateEmailCode", func(c *gin.Context) {

		token := c.DefaultPostForm("token", "nil")

		if token == "nil" {
			c.String(http.StatusNotAcceptable, fmt.Sprintln("network"))
			return
		}

		success, user := Verification.GetUserFromToken(token, err, db, router)
		if success {
			if service.SendActivateCode(user.ID, db) {
				c.JSON(200, utils.StandardResponse{
					IsSuccess: true,
					ErrorMsg:  "200",
				})
			} else {
				c.JSON(200, utils.StandardResponse{
					IsSuccess: false,
					ErrorMsg:  "0",
				})
			}

		} else {
			c.JSON(200, utils.StandardResponse{
				IsSuccess: false,
				ErrorMsg:  "501",
			})
		}

	})

	//for test
	router.GET("/frontEnd/sayHello", func(c *gin.Context) {

		// strToken := c.DefaultQuery("token", "nil")

		// success, user := Verification.GetUserFromToken(strToken, err, db, router)
		// if success {
		// 	c.String(http.StatusOK, "Hello,"+user.Email)
		// } else {
		// 	c.String(http.StatusOK, "Login expired.")
		// }

		// sid, addr := Service.GetNextUseableSlaveServer(db)
		// c.String(http.StatusOK, strconv.Itoa(sid)+",add:"+addr)

		// uuid := c.DefaultQuery("token", "nil")
		// var taskList []utils.TaskListRow = service.GetUsersTransactionList(db, uuid, 1)
		// println(len(taskList))
		// for _, value := range taskList {
		// 	println(value.Title + " " + value.Avatar + " " + value.Progress + " " + value.Status)
		// }
		// c.JSON(200, utils.TransactionListResponse{
		// 	IsSuccess: "true",
		// 	ErrorMsg:  "200",
		// 	TaskList:  taskList,
		// })
		// c.IndentedJSON(200, taskList)
	})
}

// func getSHA256HashCode(message []byte) string {
// 	hash := sha256.New()
// 	hash.Write(message)
// 	bytes := hash.Sum(nil)
// 	hashCode := hex.EncodeToString(bytes)
// 	return hashCode
// }
