package forum

import (
	"net/mail"
)

func CheckMail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
