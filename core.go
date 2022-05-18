package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Message string `json:"message"`
	FROM    string `json:"from"`
	METHOD  string `json:"method"`
	PATH    string `json:"path"`
}

func main() {
	app := fiber.New()
	hostname, err := os.Hostname()

	if err != nil {
		hostname = "Unknown Mom"
	}

	app.Use(func(c *fiber.Ctx) error {
		return c.JSON(Message{
			Message: "Hi, Mom!",
			FROM:    hostname,
			METHOD:  c.Method(),
			PATH:    c.Path(),
		})
	})

	app.Listen(":3000")
}
