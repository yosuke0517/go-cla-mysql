package main

import (
	"fmt"
	"github.com/labstack/echo"
)

func main() {
	fmt.Println("start")

	e := echo.New()
	e.Logger.Fatal(e.Start(":8080"))
}
