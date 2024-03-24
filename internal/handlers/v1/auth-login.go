package v1handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/ther0y/xeep-api/internal/services"
	"net/http"
)

func AuthLogin(c echo.Context) error {
	authService := services.NewAuthService()

	resp, err := authService.Login(c.Request().Context(), "masoods", "test")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
