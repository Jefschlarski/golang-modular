package validate

import (
	"errors"
	"unicode"
)

func Password(password string) error {
	var hasUpper, hasLower, hasNumber, hasSpecial bool
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return errors.New("password must have at least one uppercase letter, one lowercase letter, one number, one special character, and be at least 8 characters long")
	}

	return nil
}
