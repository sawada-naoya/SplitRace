//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/sawada-naoya/splitrace/handler"
	"github.com/sawada-naoya/splitrace/service"
	"github.com/sawada-naoya/splitrace/usecase"
)

// InitializeTaskHandler は handler層を起点にした依存関係の注入を行う関数。
// 以下の順で依存関係が組み立てられる：
//     TaskHandler ← TaskUsecase ← TaskService
// ・interfaceは上位層から下位層に向けて定義される（依存方向は上→下）
// ・structの生成はWireが担当（New関数で実装インスタンスを生成）
// ・handlerはusecaseのinterfaceを受け取り、usecaseはserviceのinterfaceを受け取る

func InitializeTaskHandler() *handler.TaskHandler {
	wire.Build(
		service.NewTaskService,
		usecase.NewTaskUsecase,
		handler.NewTaskHandler,
	)
	return nil
}