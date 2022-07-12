package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello worlds")
	})
	app.Get("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Hello world with values" + c.Params("id"))
	})

	app.Listen(":3000")
}
