//go:build !release
// +build !release

package main

import (
	router "test-githubaction/router"

	"github.com/gofiber/fiber/v2"
)

// 監聽 Http
func httpListen() {
	// 設定路由
	application := fiber.New()
	router.SetupRoutes(application)

	// 監聽 3001 Port
	_ = application.Listen("127.0.0.1:3001")
}
