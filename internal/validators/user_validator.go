package validators

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	isValidPhone    = regexp.MustCompile(`(\+98|0|98)?(9\d{9})`).MatchString
)

func (cv *CustomValidator) UniqueIdentifier(fl validator.FieldLevel) bool {
	identifier := fl.Field().String()

	ok, err := cv.AuthService.IsUniqueUserIdentifier(context.Background(), identifier)
	if err != nil {
		// TODO: make such error logs more readable/trackable as such errors are not being sent to the client
		//  and are only visible in the server logs
		fmt.Println(err)
		return false
	}

	return ok
}

func ValidateEmail(email string) error {
	//if err := ValidateString(email, 5, 50); err != nil {
	//	return err
	//}

	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf("is not a valid email address")
	}

	return nil
}

func ValidatePassword(password string) error {
	//if err := ValidateString(password, 6, 50); err != nil {
	//	return err
	//}

	return nil
}

func ValidatePhone(phone string) error {
	//if err := ValidateString(phone, 10, 15); err != nil {
	//	return err
	//}

	if !isValidPhone(phone) {
		return fmt.Errorf("is not a valid phone number")
	}

	return nil
}

func ValidateUsername(username string) error {
	if !isValidUsername(username) {
		return fmt.Errorf("can only contain letters, numbers and underscores")
	}

	return nil
}

func ValidateOtp(otp string) error {
	//if err := ValidateString(otp, 6, 6); err != nil {
	//	return err
	//}

	return nil
}
