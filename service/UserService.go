package service

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"

	Entity "gatelligance/entity"
	Utils "gatelligance/utils"
	Verification "gatelligance/verification"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func HandleUserLogin(password string, email string, db *gorm.DB, c *gin.Context) {
	//sha256 check
	passwdBYTE := []byte(password)
	passwdSHA := getSHA256HashCode(passwdBYTE)

	var uu = new(Entity.User)
	db.Find(&uu, "email=?", email)

	if uu.PassSHA != passwdSHA {
		//not correct passwd!
		// c.String(http.StatusOK, "passwd")
		c.JSON(http.StatusOK, Utils.LoginResponse{
			Token:     "-1",
			IsSuccess: false,
			ErrorMsg:  "402",
		})
		return
	}

	claims := &Verification.JWTClaims{
		ID: uu.ID,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(Verification.ExpireTime)).Unix()
	signedToken, err := Verification.GetToken(claims)
	if err != nil {
		// c.String(http.StatusNotFound, err.Error())
		c.JSON(http.StatusOK, Utils.LoginResponse{
			Token:     "-1",
			IsSuccess: false,
			ErrorMsg:  "401",
		})
		return
	}
	// c.String(http.StatusOK, signedToken)

	c.JSON(http.StatusOK, Utils.LoginResponse{
		Token:     signedToken,
		IsSuccess: true,
		ErrorMsg:  "200",
	})

}

func HandleUserRegister(password string, email string, nickName string, db *gorm.DB, err *error, c *gin.Context) {
	//sha256 check
	passwdBYTE := []byte(password)
	passwdSHA := getSHA256HashCode(passwdBYTE)
	//fmt.Println(passwdSHA)

	userUid := uuid.Must(uuid.NewV4(), *err)

	var uu = new(Entity.User)
	db.Find(&uu, "email=?", email)
	//fmt.Printf("phone:%s", uu.Phone)
	if uu.Email != email {
		//email num haven't used yet
		strUsrUid := userUid.String()
		newUser := Entity.User{ID: strUsrUid, Phone: "", NickName: nickName, Email: email, PassSHA: passwdSHA, Gender: "保密", Avatar: 1}
		db.Create(newUser)
		//register success
		// c.String(http.StatusOK, "ok")
		c.JSON(http.StatusOK, Utils.RegisterResponse{
			IsSuccess: true,
			ErrorMsg:  "200",
		})
	} else {
		//email is used!
		// c.String(http.StatusOK, "email")
		c.JSON(http.StatusOK, Utils.RegisterResponse{
			IsSuccess: false,
			ErrorMsg:  "301",
		})
	}

}

func getSHA256HashCode(message []byte) string {
	hash := sha256.New()
	hash.Write(message)
	bytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}
