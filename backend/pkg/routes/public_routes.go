package routes

import (
	"carbon-tax-ledger/app/controllers"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api")

	route.Post("/login", controllers.Login)
	route.Post("/logout", controllers.Logout)

	route.Get("/carbon", controllers.GetCarbonToken)
	route.Post("/pay", controllers.PayCarbonTax)
	route.Get("/history", controllers.GetCarbonPaymentHistory)

	route.Get("/wallet", controllers.GetWalletToken)
	route.Post("/topup", controllers.MintWalletToken)
}
