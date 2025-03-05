package test

import (
	"errors"
	"regexp"
)

func ValidateEmail(email string) (bool, error) {
	if email == "" {
		return false, errors.New("email cannot be empty")
	}
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(regex, email)
	return match, nil
}

//Test for valid email formats.Test for invalid email formats (e.g., missing @, no domain).Test for an empty email input.
