package email

import (
	"online_chat/enviroment"
	gomail "gopkg.in/mail.v2"
)

var (
	email = enviroment.GoDotEnvVariable("EMAIL")
	email_smtp = enviroment.GoDotEnvVariable("EMAIL_SMTP")
	email_password = enviroment.GoDotEnvVariable("EMAIL_PASSWORD")
)
func EmailSender(user_email string, code string) error {
	e := gomail.NewMessage()
	e.SetHeader("From", email)
	e.SetHeader("To", user_email)
	e.SetHeader("Subject", "Ссылка для восстановления пароля")
	e.SetBody("text/plain", "https://tt-chat.danyatochka.ru/recovery?token="+code)

	dialer := gomail.NewDialer(email_smtp, 587, email, email_password)

	return dialer.DialAndSend(e)
}