package forum

import (
	"errors"
	"regexp"
	"strings"
)

type EmailStorage struct {
	Email string
}

func (e *EmailStorage) IsValidEmail() bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(e.Email)
}

func (e *EmailStorage) FormatEmail() {
	e.Email = strings.ToLower(strings.TrimSpace(e.Email))
}

func (e *EmailStorage) SetEmail(email string) error {
	if !e.IsValidEmail() {
		return errors.New("adresse e-mail invalide")
	}
	e.Email = email
	return nil
}
