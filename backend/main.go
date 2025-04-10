package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const (
	apiUrl = ""
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API de To-Do List funcionando!")
	})
	fmt.Println("API de To-Do List funcionando!")
	app.Listen(":7070")
}
