package main

import (
	"regexp"

	gomail "gopkg.in/gomail.v2"
)

// ValidEmail return whether the email is valid
func ValidEmail(s string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(s)
}

// SendMail 发送邮件
func SendMail(tl, bd string, recvs []string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", "1@qq.com")
	m.SetHeader("To", recvs...)
	m.SetHeader("Subject", tl)
	m.SetBody("text/plain", bd)

	d := gomail.NewDialer("conf.Config.Smtp.Host", "conf.Config.Smtp.Port", "conf.Config.Smtp.SmtpUser", "conf.Config.Smtp.SmtpPassword")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	return nil
}

func main() {
	e := []string{"2@qq.com", "3@qq.com"}
	SendMail("title", "body", e)
}
