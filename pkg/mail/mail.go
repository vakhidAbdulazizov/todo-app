package mail

import (
	"net/smtp"
	"os"
)

func SendEmail(email string, message []byte) error {
	auth := smtp.PlainAuth("", os.Getenv("EMAIL_FROM"), os.Getenv("EMAIL_PASSWORD"), os.Getenv("EMAIL_SMT_HOST"))

	err := smtp.SendMail(os.Getenv("EMAIL_SMT_HOST")+":"+os.Getenv("EMAIL_SMT_PORT"), auth, os.Getenv("EMAIL_FROM"), []string{email}, message)

	if err != nil {
		return err
	}
	return nil
}
