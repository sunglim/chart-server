// A server providing rest apis to receive operatoins or statistics data.
package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New()
	app.Get("/insert", InsertPriceHandler)

	log.Fatal(app.Listen(":3000"))
}

func InsertPriceHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
