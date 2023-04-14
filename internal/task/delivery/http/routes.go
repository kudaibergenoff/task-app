package http

import (
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	taskHandler := NewTaskHandler(e)
	g := e.Group("/task")
	g.GET("/get-all", taskHandler.GetAll)
}
