package forum

import (
	"unicode"
)

func CheckPassword(testpassword string) bool {
	var (
		hasUpperCase bool
		hasLowerCase bool
		hasNumber    bool
	)

	for _, char := range testpassword {
		switch {
		case unicode.IsUpper(char):
			hasUpperCase = true
		case unicode.IsLower(char):
			hasLowerCase = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}

	if !hasUpperCase {
		return false
	}
	if !hasLowerCase {
		return false
	}
	if !hasNumber {
		return false
	}
	if len(testpassword) < 8 {
		return false
	}

	return true
}
