package v1handlers

import "github.com/labstack/echo/v4"

func SetupRoutes(e *echo.Echo) {
	v1Routes := e.Group("/v1")
	v1Routes.POST("/auth/login", authLogin)
}
