//go:build release
// +build release

package main

import (
	router "test-githubaction/router"
	
	"github.com/gofiber/fiber/v2"
)

// 監聽 Http 
func httpListen() {
	// 設定路由
	application := fiber.New(fiber.Config{Prefork: true})
	router.SetupRoutes(application)

	// 監聽 3000 Port
	_ = application.Listen(":3001")
}
