package routes

import (
	"coding-chelleng/core/presentation/http/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, c handlers.CarHandler, o handlers.OrderHandler) {
	api := app.Group("/api/v1")

	car := api.Group("/cars")
	car.Post("/", c.AddCar)
	car.Get("/:id", c.GetCarById)
	car.Patch("/:id", c.UpdateCar)
	car.Delete("/:id", c.DeleteCar)
	car.Get("/", c.GetCars)

	order := api.Group("/order")
	order.Post("/", o.AddOrder)
	order.Delete("/:id", o.DeleteOrder)
	order.Get("/", o.GetOrders)
}
