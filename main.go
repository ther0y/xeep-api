package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	v1handlers "github.com/ther0y/xeep-api/internal/handlers/v1"
	"github.com/ther0y/xeep-api/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	appPort, err := utils.GetEnv("SERVER_PORT")
	if err != nil || appPort == "" {
		appPort = "1323"
		fmt.Println("SERVER_PORT not found in .env file, using default port 1323")
	}

	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	v1Routes := e.Group("/v1")

	// Auth routes
	v1Routes.POST("/auth/login", v1handlers.AuthLogin)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	e.Logger.Fatal(e.Start(":" + appPort))
}
