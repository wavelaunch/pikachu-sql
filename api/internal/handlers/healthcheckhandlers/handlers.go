package healthcheckhandlers

import "github.com/gofiber/fiber/v2"

type handler struct {
}

func New() *handler {
	return &handler{}
}

func (h *handler) HealthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}
