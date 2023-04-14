package http

import (
	"github.com/kudaibergenoff/task-app/internal/task/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TaskHandler struct {
	taskRepo repository.TaskRepository
	e        *echo.Echo
}

func NewTaskHandler(e *echo.Echo, taskRepo repository.TaskRepository) *TaskHandler {
	return &TaskHandler{
		taskRepo: taskRepo,
		e:        e,
	}
}

func (handler *TaskHandler) Index(c echo.Context) error {
	return c.String(http.StatusOK, "Index page")
}

func (handler *TaskHandler) GetAll(ctx echo.Context) error {
	tasks, err := handler.taskRepo.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve tasks",
		})
	}
	return ctx.JSON(http.StatusOK, tasks)
}
