package controllers

import (
	"carbon-tax-ledger/app/models"

	"github.com/gofiber/fiber/v2"
)

func PayCarbonTax(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Carbon tax paid successfully",
		"data": fiber.Map{
			"amount": 100,
		},
	})
}

func GetCarbonPaymentHistory(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Carbon tax payment history retrieved successfully",
		"data": map[string][]models.Payment{
			"payments": {
				{
					ID:     "1",
					Amount: 100,
				},
				{
					ID:     "2",
					Amount: 200,
				},
			},
		},
	})
}

func GetCarbonToken(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Info retrieved successfully",
		"data": fiber.Map{
			"token": 69,
		},
	})
}
