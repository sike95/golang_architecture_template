package controllers

import "github.com/gofiber/fiber/v2"

var (
	HealthzController healthzControllerInterface = &healthzController{}
)

type healthzControllerInterface interface {
	Check(c *fiber.Ctx) error
}

type healthzController struct{}

func (h *healthzController) Check(c *fiber.Ctx) error {
	return c.Status(200).SendString("Connected")
}
