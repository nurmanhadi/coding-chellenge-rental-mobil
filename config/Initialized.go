package config

import (
	"coding-chelleng/core/application/services"
	"coding-chelleng/core/entities/repository"
	"coding-chelleng/core/infrastructure/database"
	"coding-chelleng/core/infrastructure/database/migration"
	"coding-chelleng/core/presentation/http/handlers"
	"coding-chelleng/core/presentation/http/routes"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Initialized(app *fiber.App) {
	db := database.ConnectDatabase()
	migration.AutoMigrate(db)
	validate := validator.New()

	//car
	cRepo := repository.NewCarRepository(db)
	cService := services.NewCarService(cRepo, validate)
	cHandler := handlers.NewCarHandler(&cService)

	//order
	oRepo := repository.NewOrderRepository(db)
	oService := services.NewOrderService(oRepo, validate)
	oHandler := handlers.NewOrderHandler(&oService)

	//route
	routes.SetupRoutes(app, cHandler, oHandler)
}
