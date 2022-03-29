package service

import (
	"fmt"
	"gatelligance/entity"
	"math/rand"
	"strings"
	"time"

	"github.com/go-gomail/gomail"
	"github.com/jinzhu/gorm"
)

type EmailParam struct {
	// ServerHost 邮箱服务器地址，如腾讯邮箱为smtp.qq.com
	ServerHost string
	// ServerPort 邮箱服务器端口，如腾讯邮箱为465
	ServerPort int
	// FromEmail　发件人邮箱地址
	FromEmail string
	// FromPasswd 发件人邮箱密码（注意，这里是明文形式），TODO：如果设置成密文？
	FromPasswd string
	// Toers 接收者邮件，如有多个，则以英文逗号(“,”)隔开，不能为空
	Toers string
	// CCers 抄送者邮件，如有多个，则以英文逗号(“,”)隔开，可以为空
	CCers string
}

// 全局变量，因为发件人账号、密码，需要在发送时才指定
// 注意，由于是小写，外面的包无法使用
var serverHost, fromEmail, fromPasswd string
var serverPort int

var m *gomail.Message

func InitEmail(ep *EmailParam) {
	toers := []string{}

	serverHost = ep.ServerHost
	serverPort = ep.ServerPort
	fromEmail = ep.FromEmail
	fromPasswd = ep.FromPasswd

	m = gomail.NewMessage()

	if len(ep.Toers) == 0 {
		return
	}

	for _, tmp := range strings.Split(ep.Toers, ",") {
		toers = append(toers, strings.TrimSpace(tmp))
	}

	// 收件人可以有多个，故用此方式
	m.SetHeader("To", toers...)

	//抄送列表
	if len(ep.CCers) != 0 {
		for _, tmp := range strings.Split(ep.CCers, ",") {
			toers = append(toers, strings.TrimSpace(tmp))
		}
		m.SetHeader("Cc", toers...)
	}

	// 发件人
	// 第三个参数为发件人别名，如"李大锤"，可以为空（此时则为邮箱名称）
	m.SetAddressHeader("From", fromEmail, "")
}

// SendEmail body支持html格式字符串
func SendEmail(subject, body string) {
	// 主题
	m.SetHeader("Subject", subject)

	// 正文
	m.SetBody("text/html", body)

	d := gomail.NewPlainDialer(serverHost, serverPort, fromEmail, fromPasswd)
	// 发送
	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
}

func SendActivateCode(uuid string, db *gorm.DB) bool {
	var uu []entity.User

	db.Find(&uu, "id=?", uuid)

	if len(uu) == 0 {
		fmt.Printf("user email send: user not found\n")
		return false
	}

	vcode := get6DigitCode()

	var aa []entity.EmailActiCode
	db.Find(&aa, "uuid=?", uuid)
	if len(aa) != 0 {
		db.Delete(aa[0])
	}

	db.Create(entity.EmailActiCode{
		Uuid: uuid,
		Code: vcode,
	})

	sendEmail(vcode, uu[0].Email)
	return true
}

func ActivateEmail(uuid string, code string, db *gorm.DB) bool {
	var uu []entity.User

	db.Find(&uu, "id=?", uuid)

	if len(uu) == 0 {
		fmt.Printf("user email activatation: user not found\n")
		return false
	}

	var aa []entity.EmailActiCode
	db.Find(&aa, "uuid=?", uuid)
	if len(aa) == 0 {
		fmt.Printf("user email activatation: code not found\n")
		return false
	}

	if aa[0].Code == code {
		db.Delete(uu[0])
	}
	uu[0].Activated = 1
	db.Create(uu[0])

	return true
}

func sendEmail(vcode string, to string) {
	serverHost := "smtp.163.com"
	serverPort := 465
	fromEmail := "jiangjm718@163.com" //发件人邮箱
	fromPasswd := "NOEQLRDWEFJLJZHC"  //授权码

	myToers := to // 收件人邮箱，逗号隔开
	myCCers := to //"readchy@163.com"

	subject := "欢迎使用凝智成林"
	// body := `您正在验证邮箱，验证码是 113740.<br>
	//          来自 <a href = "http://c.biancheng.net/"> 凝智成林Gatelligance</a>`
	body := `您正在验证邮箱，验证码是 ` + vcode + ` .<br>` +
		`来自 <a href = "https://j1am1ng.github.io/Gatelligence-pc"> 凝智成林Gatelligance</a>`
	// 结构体赋值
	myEmail := &EmailParam{
		ServerHost: serverHost,
		ServerPort: serverPort,
		FromEmail:  fromEmail,
		FromPasswd: fromPasswd,
		Toers:      myToers,
		CCers:      myCCers,
	}

	InitEmail(myEmail)
	SendEmail(subject, body)
}

func get6DigitCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vcode
}
