package services

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/k3a/html2text"
	"github.com/vanng822/go-premailer/premailer"
	"gopkg.in/gomail.v2"
	"html/template"
	"leizhenpeng/go-email-verification/config"
	"os"
	"path/filepath"
)

type EmailData struct {
	URL      string
	UserName string
	Subject  string
}

func GenEmailVerifyURL(code string) string {
	configNow := config.GetConfig()
	return fmt.Sprintf("%s/verify?code=%s", configNow.BaseUrl, code)
}

func GenEmailData(email string, code string) *EmailData {
	return &EmailData{
		URL:      GenEmailVerifyURL(code),
		UserName: email,
		Subject:  "请激活您的账号",
	}
}

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendEmail(email string, data *EmailData) error {

	// Sender data.
	configNow := config.GetConfig()
	from := configNow.EmailFrom
	smtpPass := configNow.SmtpPass
	smtpUser := configNow.SmtpUser
	to := email
	smtpHost := configNow.SmtpHost
	smtpPort := configNow.SmtpPort

	var body bytes.Buffer

	template, err := ParseTemplateDir("templates")
	if err != nil {
		return errors.New("could not parse template")
	}

	template.ExecuteTemplate(&body, "email-verify.html", &data)
	htmlString := body.String()
	prem, _ := premailer.NewPremailerFromString(htmlString, nil)
	htmlInline, err := prem.Transform()
	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", htmlInline)
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		return errors.New("could not send email")
	}
	return nil

}
