package routes

import (
	"elisa-custom-hash-srv/app/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// PublicRoutes func for describe group of public routes.
func Routes(app *fiber.App) {
	// Create routes group.
	route := app.Group("connect-custom-hash-srv")
	app.Use(cors.New())
	// Routes for GET method:
	route.Get("ping", controllers.Ping)
	route.Post("/", controllers.Process)
}
