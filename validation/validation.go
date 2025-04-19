package validation

import (
	"regexp"
)


func IsValidPhoneNumber(phoneNumber string) bool{
	//allow only  10 digits, spaces, and dashes.
	phoneNumberPattern := `^\d{3}[-\s]?\d{3}[-\s]?\d{4}$`

	matched, err := regexp.MatchString(phoneNumberPattern, phoneNumber)

	if err != nil {
		return false
	}

	return matched;
}


func IsValidEmail(email string) bool {
	emailPattern := `[a-zA-z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`;
	matched, err := regexp.MatchString(emailPattern, email)
	if err != nil {
		return false
	}
	return matched
}