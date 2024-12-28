package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func MintWalletToken(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Top up successful",
		"data": fiber.Map{
			"token": 999,
		},
	})
}

func GetWalletToken(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Wallet token retrieved successfully",
		"data": fiber.Map{
			"token": 999,
		},
	})
}
