package usecase

import (
	"github.com/sawada-naoya/splitrace/dto"
	"github.com/sawada-naoya/splitrace/service"
)

// 上層が期待する動き（interface）
type TaskUsecase interface {
	RunSerialAndParallel(count int) dto.RunTaskResponse
}

// 実態（中身は TaskService を注入して使う）
type taskUsecase struct {
	ts service.TaskService
}

// DIのための生成関数（Wireがここを呼ぶ）
func NewTaskUsecase(ts service.TaskService) TaskUsecase {
	return &taskUsecase{
		ts: ts,
	}
}

func (u *taskUsecase) RunSerialAndParallel(count int) dto.RunTaskResponse {
	serial := u.ts.RunSerial(count)
	parallel := u.ts.RunParallel(count)

	return dto.RunTaskResponse{
		Count: count,
		Serial: serial,
		Parallel: parallel,
	}
}