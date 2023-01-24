package utils

import (
	"fmt"
	"log"
	"bytes"
	"net/smtp"
	"text/template"

	"app/config"
)


func Send(signupPassword, email string) {

	from := config.SMTP_EMAIL
	password := config.SMPT_PASSWORD
	to := []string{ email }
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles("index.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	body.Write([]byte(fmt.Sprintf("Subject: This is a tes \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Message string
	}{
		Message: "Your Password:" + signupPassword,
	})

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	
	if err != nil { log.Println(err); return }

	log.Println("Send Email password")
}