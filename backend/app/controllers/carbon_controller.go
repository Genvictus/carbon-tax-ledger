package controllers

import (
	"bytes"
	"carbon-tax-ledger/app/models"
	"carbon-tax-ledger/contract"
	"carbon-tax-ledger/pkg/repository"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func PayCarbonTax(c *fiber.Ctx) error {
	req := &models.PayRequest{}
	if err := c.BodyParser(req); err != nil {
		return handleErrorResponse(c, fiber.ErrBadRequest.Code, "Failed to get request", err)
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return handleErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid request body", err)
	}

	contract, err := contract.GetContract(c, repository.ChainCodeName["CT"])
	if err != nil {
		return handleErrorResponse(c, fiber.ErrUnauthorized.Code, "Failed to get contract", err)
	}

	_, err = contract.SubmitTransaction("Pay", strconv.Itoa(req.Amount), "")
	if err != nil {
		return handleErrorResponse(c, fiber.ErrBadRequest.Code, "Failed to pay carbon tax", err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Carbon tax paid successfully",
		"data": fiber.Map{
			"amount": req.Amount,
		},
	})
}

func GetCarbonToken(c *fiber.Ctx) error {
	contract, err := contract.GetContract(c, repository.ChainCodeName["CT"])
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
		"message": "Info retrieved successfully",
		"data": fiber.Map{
			"token": resultInt,
		},
	})
}

// Format JSON data
func formatJSON(data []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
		return "", fmt.Errorf("failed to format JSON data: %v", err)
	}
	return prettyJSON.String(), nil
}
