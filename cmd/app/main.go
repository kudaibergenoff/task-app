package main

import (
	"github.com/kudaibergenoff/task-app/internal/task/delivery/http"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	http.SetupRoutes(e)

	err := e.Start(":8080")
	if err != nil {
		return
	}
}
