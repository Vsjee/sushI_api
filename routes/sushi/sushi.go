package sushi

import (
	"sushi-api/routes/sushi/controllers"

	"github.com/gofiber/fiber/v2"
)

func SushiRoute(app *fiber.App) {
	app.Get("/sushi", controllers.GetAllSushi)
}
