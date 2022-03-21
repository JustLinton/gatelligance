package service

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"

	Entity "gatelligance/entity"
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

	if uu.PassSHA == passwdSHA {
		//passwd is correct
		//login success
		// c.SetCookie("user_token", uu.ID, 1000, "/", domain, false, true)
		c.String(http.StatusOK, "ok\n")

	} else {
		//not correct passwd!
		c.String(http.StatusOK, "passwd")
		return
	}

	claims := &Verification.JWTClaims{
		ID: uu.ID,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(Verification.ExpireTime)).Unix()
	signedToken, err := Verification.GetToken(claims)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, signedToken)
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
		newUser := Entity.User{ID: strUsrUid, Phone: "", NickName: nickName, Email: email, PassSHA: passwdSHA, Gender: "保密"}
		db.Create(newUser)
		//register success
		c.String(http.StatusOK, "ok")
	} else {
		//email is used!
		c.String(http.StatusOK, "email")
	}
}

func getSHA256HashCode(message []byte) string {
	hash := sha256.New()
	hash.Write(message)
	bytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}
