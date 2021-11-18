package router

import (
	"github.com/gofiber/fiber/v2"
)

// 設定路由
func SetupRoutes(application *fiber.App) {
	// 將靜態頁面綁定根目錄
	application.Static("/", "./dist")
}
