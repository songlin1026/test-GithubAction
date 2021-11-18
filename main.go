package main

import (
	router "test-githubaction/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.SetupRoutes(app)

	app.Get("/root", func(c *fiber.Ctx) error {
		return c.SendString("root")
	})

	_ = app.Listen(":3001")
}
