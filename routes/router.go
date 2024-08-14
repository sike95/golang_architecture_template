package routes

import (
	"golang_architectire_template/controllers"

	"github.com/gofiber/fiber/v2"
)

func CreateRoutes(router fiber.Router) {
	router.Get("/healthz", controllers.HealthzController.Check)

	router.Get("/user", controllers.Create)
	router.Get("/user", controllers.Get)

}
