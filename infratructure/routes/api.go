package routes

import (
	"github.com/labstack/echo"
	"go-cla-mysql/adapters/controllers"
	"go-cla-mysql/injector"
)

func InitRoutign(e *echo.Echo) {
	todo := injector.InjectTodo()
	e.GET("/health", controllers.HealthCheckController())
	g := e.Group("/api/v1")
	{
		g.GET("/hello", controllers.HelloController())
		g.GET("/todo/get", todo.GetAll())
	}
}
