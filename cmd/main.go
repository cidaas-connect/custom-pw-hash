package main

import (
	"fmt"

	"elisa-custom-hash-srv/pkg/configs"
	"elisa-custom-hash-srv/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Routes.
	routes.Routes(app)

	app.Listen(":3000")

	fmt.Println("Service Started!")

}
