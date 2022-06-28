package test

import (
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestSendEamil(t *testing.T) {
	m := gomail.NewMessage()
	// 发送人
	m.SetHeader("From", "1037134254@qq.com")
	// 收件人多个/一个
	m.SetHeader("To", "2865549101@qq.com")
	// 抄送
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	// 标题
	m.SetHeader("Subject", "Hello!")
	n := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000)
	s := strconv.FormatInt(n, 10)
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
}
