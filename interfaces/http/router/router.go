package router

import (
	"sandbox/interfaces/http/handler"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App, transactionHandler *handler.TransactionHandler) {
	api := app.Group("/api")

	// Transaction routes
	api.Post("/upload", transactionHandler.UploadAndExtract)
	api.Post("/upload/detailed", transactionHandler.UploadAndExtractDetailed)
	api.Post("/report/excel", transactionHandler.GenerateRecapExcel)

	// Health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "healthy",
		})
	})
}
