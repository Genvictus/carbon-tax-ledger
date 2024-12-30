package controllers

import (
	"carbon-tax-ledger/app/models"
	"carbon-tax-ledger/contract"
	"carbon-tax-ledger/pkg/repository"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func MintWalletToken(c *fiber.Ctx) error {
	req := &models.TopUpRequest{}
	if err := c.BodyParser(req); err != nil {
		return handleErrorResponse(c, fiber.ErrBadRequest.Code, "Failed to get request", err)
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return handleErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid request body", err)
	}

	contract, err := contract.GetContract(c, repository.ChainCodeName["WT"])
	if err != nil {
		return handleErrorResponse(c, fiber.ErrUnauthorized.Code, "Failed to get contract", err)
	}

	_, err = contract.SubmitTransaction("Mint", strconv.Itoa(req.Amount), "")
	if err != nil {
		return handleErrorResponse(c, fiber.ErrBadRequest.Code, "Failed to mint token", err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Top up successful",
		"data": fiber.Map{
			"token": req.Amount,
		},
	})
}

func GetWalletToken(c *fiber.Ctx) error {
	contract, err := contract.GetContract(c, repository.ChainCodeName["WT"])
	if err != nil {
		return handleErrorResponse(c, fiber.ErrUnauthorized.Code, "Failed to get contract", err)
	}

	evaluateResult, err := contract.EvaluateTransaction("ClientAccountBalance")
	if err != nil {
		return handleErrorResponse(c, fiber.ErrBadRequest.Code, "Failed to evaluate transaction", err)
	}
	result, err := formatJSON(evaluateResult)
	if err != nil {
		return handleErrorResponse(c, fiber.ErrBadRequest.Code, "Failed to format JSON data", err)
	}

	resultInt, _ := strconv.Atoi(result)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Wallet token retrieved successfully",
		"data": fiber.Map{
			"token": resultInt,
		},
	})
}
