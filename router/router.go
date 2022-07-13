package router

import (
	"github.com/erfandiakoo/TimeApi/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Worked")
	})

	time := api.Group("api")
	time.Get("/now", handler.GetTime)
	time.Get("/unix", handler.GetTimeUnix)
	time.Get("/utc", handler.GetTimeUtc)
	time.Get("/ip", handler.GetTimeIP)
	time.Get("timezone/:area", handler.GetTimeTimeZone)

}
