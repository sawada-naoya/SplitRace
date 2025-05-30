package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/sawada-naoya/splitrace/di"
	"github.com/sawada-naoya/splitrace/router"
)

func main() {
	e := echo.New()

	// CORSミドルウェアの設定（ローカルのNext.jsからのリクエストを許可）
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
	}))

	// DI経由でhandlerを注入
	taskHandler := di.InitializeTaskHandler()

	// ルーティング
	router.InitRouter(e, taskHandler)

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}