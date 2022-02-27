package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wavelaunch/pikachu-sql/api/internal/app"
)

func registerRoutes(router *fiber.App, app *app.Application) {
	router.Get("/", app.Handlers.HealthCheckHandlers.HealthCheck)

	// v1 := router.Group("/v1")
	// {
	// 	v1.Post("/task", app.Handlers.TaskHandlers.CreateTask)
	// }
}
