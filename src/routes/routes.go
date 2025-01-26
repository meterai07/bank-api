package routes

import (
	"bank/src/handlers"
	"bank/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, nasabahHandler *handlers.NasabahHandler, authHandler *handlers.AuthHandler) {
	api := app.Group("/api/v1")

	// Public routes
	api.Post("/daftar", nasabahHandler.Daftar)
	api.Post("/login", authHandler.Login)

	// Protected routes
	protected := api.Use(middleware.AuthRequired)
	protected.Post("/tabung", nasabahHandler.Tabung)
	protected.Post("/tarik", nasabahHandler.Tarik)
	protected.Get("/saldo/:no_rekening", nasabahHandler.Saldo)
}
