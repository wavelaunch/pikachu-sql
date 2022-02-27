package ports

import "github.com/gofiber/fiber/v2"

type HealthCheckHandlers interface {
	HealthCheck(c *fiber.Ctx) error
}

type TaskHandlers interface {
}
