package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"

	Service "gatelligance/service"
	Verification "gatelligance/verification"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitUsersController(err *error, db *gorm.DB, router *gin.Engine) {

	router.POST("/frontEnd/login", func(c *gin.Context) {

		password := c.DefaultPostForm("password", "nil")
		email := c.DefaultPostForm("email", "nil")

		if password == "nil" || email == "nil" {
			c.String(http.StatusNotAcceptable, fmt.Sprintln("network"))
			return
		}

		Service.HandleUserLogin(password, email, db, c)
	})

	router.POST("/frontEnd/register", func(c *gin.Context) {

		email := c.DefaultPostForm("email", "nil")
		password := c.DefaultPostForm("password", "nil")
		nickName := c.DefaultPostForm("nickName", "nil")

		if email == "nil" || password == "nil" || nickName == "nil" {
			c.String(http.StatusNotAcceptable, fmt.Sprintln("network"))
		}

		Service.HandleUserRegister(password, email, nickName, db, err, c)
	})

	router.GET("/frontEnd/refreshToken", Verification.RefreshTokenHandler)

	//for test
	router.GET("/frontEnd/sayHello", func(c *gin.Context) {

		strToken := c.DefaultQuery("token", "nil")
		claim, stat := Verification.VerifyToken(strToken)
		if !stat {
			c.String(http.StatusOK, "Login expired.")
			return
		}
		c.String(http.StatusOK, "Hello,"+claim.ID)
	})
}

func getSHA256HashCode(message []byte) string {
	hash := sha256.New()
	hash.Write(message)
	bytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}
