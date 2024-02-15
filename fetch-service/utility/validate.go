package utility

import "regexp"

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`)
	return emailRegex.MatchString(email)
}
