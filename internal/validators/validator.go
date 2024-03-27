package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/ther0y/xeep-api/internal/services"
	"net/http"
)

// ValidationErrorObject represents a single field's validation error
type ValidationErrorObject struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// CustomValidationError wraps the slice of ValidationErrorObject
type CustomValidationError struct {
	Errors []ValidationErrorObject
}

// Error makes CustomValidationError satisfy the error interface
func (cve CustomValidationError) Error() string {
	return "validation failed"
}

type CustomValidator struct {
	Validator   *validator.Validate
	HttpClient  *http.Client
	AuthService services.AuthService
}

func (cv *CustomValidator) RegisterCustomValidation() {
	_ = cv.Validator.RegisterValidation("uniqueIdentifier", cv.UniqueIdentifier)
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Convert the validation errors into a slice of ValidationErrorObject
		validationErrors := cv.TranslateErrors(err)
		// Return an echo.HTTPError with the structured validation errors
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"errors": validationErrors,
		})
	}
	return nil
}

func (cv *CustomValidator) TranslateErrors(err error) []ValidationErrorObject {
	var errors []ValidationErrorObject
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errors = append(errors, ValidationErrorObject{
				Field:   e.Field(),
				Message: e.Tag(),
			})
		}
	}
	return errors
}

//func (cv *CustomValidator) translateError(e validator.FieldError) string {
//	switch e.Tag() {
//	case "required":
//		return e.Field() + " is required"
//	case "uniqueIdentifier":
//		return e.Field() + " is not unique"
//	default:
//		return e.Field() + " is invalid"
//	}
//}
