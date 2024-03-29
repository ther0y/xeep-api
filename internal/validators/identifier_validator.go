package validators

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
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
