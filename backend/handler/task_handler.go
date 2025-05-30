package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawada-naoya/splitrace/dto"
	"github.com/sawada-naoya/splitrace/usecase"
)


type TaskHandler struct {
	ts usecase.TaskUsecase
}

func NewTaskHandler(ts usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{
		ts: ts,
	}
}

func (h *TaskHandler) RunTasks(c echo.Context) error {
	var req dto.RunTasckRequest
	if err := c.Bind(&req); err != nil || req.Count <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	res := h.ts.RunSerialAndParallel(req.Count)
	return c.JSON(http.StatusOK, res)
}