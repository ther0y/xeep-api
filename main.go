package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/ther0y/xeep-api/internal/handlers/v1"
	"github.com/ther0y/xeep-api/internal/services"
	"github.com/ther0y/xeep-api/internal/utils"
	"github.com/ther0y/xeep-api/internal/validators"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	appPort, err := utils.GetEnv("SERVER_PORT")
	if err != nil || appPort == "" {
		appPort = "1323"
		fmt.Println("SERVER_PORT not found in .env file, using default port 1323")
	}

	authService := services.NewAuthService()
	newValidator := validator.New()
	httpClient := &http.Client{}

	customValidator := &validators.CustomValidator{
		Validator:   newValidator,
		AuthService: authService,
		HttpClient:  httpClient,
	}
	customValidator.RegisterCustomValidation()

	e := echo.New()
	e.Validator = customValidator

	// Custom HTTP error handler
	e.HTTPErrorHandler = func(err error, c echo.Context) {

		if he, ok := err.(*echo.HTTPError); ok {
			if errMap, ok := he.Message.(map[string]interface{}); ok {
				if errors, ok := errMap["errors"]; ok {
					// Send structured error messages as JSON
					c.JSON(he.Code, map[string]interface{}{
						"errors": errors,
					})
					return
				}
			}
			// If not a map or does not contain "errors", send the original message
			c.JSON(he.Code, map[string]interface{}{
				"message": he.Message,
			})
			return
		}

		// Fallback to default error handler for other types of errors
		e.DefaultHTTPErrorHandler(err, c)
	}

	v1Routes := e.Group("/v1")

	// Auth routes
	v1Routes.POST("/auth/login", v1handlers.AuthLogin)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	e.GET("/test-error", func(c echo.Context) error {
		return echo.NewHTTPError(http.StatusInternalServerError, "Intentional error")
	})

	e.Logger.Fatal(e.Start(":" + appPort))
}
