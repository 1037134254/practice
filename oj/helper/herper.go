package helper

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	IsAdmin  int    `json:"isAdmin"`
	jwt.StandardClaims
}

// 生成MD5
func MD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

var tokenKey = []byte("gin-oj-key")

// 生成token
func GenerateToken(identity, name string, isAdmin int) (string, error) {
	UserClaim := &UserClaims{
		Identity:       identity,
		Name:           name,
		IsAdmin:        isAdmin,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(tokenKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyseToken 解析token
func AnalyseToken(tokens string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokens, userClaim, func(token *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("AnalyseToken Error：%v", err)
	}

	return userClaim, nil
}

// 验证码
func SendCode(email string, code int64) error {
	m := gomail.NewMessage()
	// 发送人
	m.SetHeader("From", "1037134254@qq.com")
	// 收件人多个/一个
	m.SetHeader("To", email)
	// 抄送
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	// 标题
	m.SetHeader("Subject", "Hello!")
	s := strconv.FormatInt(code, 10)
	// 主题
	m.SetBody("text/html", "Hello <b>你的验证码是</b>"+s)
	// 附件
	//m.Attach("C:\\oj\\go.mod")
	d := gomail.NewDialer("smtp.qq.com", 25, "1037134254@qq.com", "xzgabsgfxsjpbcfd")
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	log.Printf("发送成功")
	return nil
}

// uuid
func UUid() string {
	return uuid.NewV4().String()
}

//生成验证码
func GetRand() int64 {
	code := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000)
	return code
}

//保存code
func CodeSave(code []byte) (string, error) {
	year := strconv.Itoa(time.Now().Year())
	month := strconv.Itoa(int(time.Now().Month()))
	day := strconv.Itoa(time.Now().Day())
	dir := "code/" + year + "-" + month + "-" + day + UUid()
	path := dir + "/main.go"
	err := os.Mkdir(dir, 0777)
	if err != nil {
		return "", nil
	}
	create, err := os.Create(path)
	if err != nil {
		return "", err
	}
	create.Write(code)
	defer create.Close()
	return path, nil
}
