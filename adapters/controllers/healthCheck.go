package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func HealthCheckController() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "I'm fine\n")
	}
}

func HelloController() echo.HandlerFunc {
	return func(c echo.Context) error {
		jsonMap := map[string]string{
			"hello": "Hello",
		}
		return c.JSON(http.StatusOK, jsonMap)
	}
}
