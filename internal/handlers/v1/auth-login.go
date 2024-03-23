package v1handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func AuthLogin(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login",
	})
}
