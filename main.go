package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}

func main() {
	fmt.Println("start")

	e := echo.New()

	// 仮ルーティング
	e.GET("/", Hello())
	e.Logger.Fatal(e.Start(":8080"))
}
