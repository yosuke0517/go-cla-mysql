package routes

import (
	"github.com/labstack/echo"
	"go-cla-mysql/adapters/controllers"
)

func InitRoutign(e *echo.Echo) {
	e.GET("/health", controllers.HealthCheckController())
	g := e.Group("/api/v1")
	{
		g.GET("/hello", controllers.HelloController())
	}
}
