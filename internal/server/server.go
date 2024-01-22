// A server providing rest apis to receive operatoins or statistics data.
package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New()
	app.Get("/insert", InsertPriceHandler)
	app.Get("/rsi", GetRsi)

	log.Fatal(app.Listen(":3000"))
}

func InsertPriceHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func GetRsi(c *fiber.Ctx) error {
	return c.SendString("rsi")
}
