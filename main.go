package main

import (
	"fmt"
	"github.com/labstack/echo"
	"go-cla-mysql/routes"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample swagger server.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /api/v1
func main() {
	fmt.Println("start")

	e := echo.New()

	// ルーティング
	routes.InitRoutign(e)
	e.Logger.Fatal(e.Start(":8080"))
}
