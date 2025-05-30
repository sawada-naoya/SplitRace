package handler

import (
	"log"
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
	var req dto.RunTaskRequest
	if err := c.Bind(&req); err != nil || req.Count <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	log.Println("受信したリクエスト:", req)
	if req.Count <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid count value",
		})
	}

	res := h.ts.RunSerialAndParallel(req.Count)
	return c.JSON(http.StatusOK, res)
}