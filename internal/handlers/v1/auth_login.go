package v1handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/ther0y/xeep-api/internal/services"
	"net/http"
)

type payload struct {
	Identifier string `json:"identifier" validate:"required,min=3,max=100"`
	Password   string `json:"password" validate:"required,min=4,max=50"`
}

func authLogin(c echo.Context) error {
	authService := services.NewAuthService()

	payload := new(payload)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid payload",
		})
	}

	if err := c.Validate(payload); err != nil {
		return err
	}

	resp, err := authService.Login(c.Request().Context(), payload.Identifier, payload.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
