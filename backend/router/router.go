package router

import (
	"github.com/sawada-naoya/splitrace/handler"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo, h *handler.TaskHandler) {
	e.POST("/run-dual-tasks", h.RunTasks)
}
