package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/wavelaunch/pikachu-sql/api/internal/app"
)

func main() {
	app := app.New()
	router := fiber.New(fiber.Config{
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	})

	registerRoutes(router, app)

	app.Log.Error.Fatal(router.Listen(":8080"))
}
