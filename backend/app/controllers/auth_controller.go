package controllers

import (
	"carbon-tax-ledger/pkg/repository"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Login handles user login with file uploads and session creation.
func Login(c *fiber.Ctx) error {
	// Validate required form values
	mspID := c.FormValue("mspID")
	if mspID == "" {
		return handleErrorResponse(c, fiber.StatusBadRequest, "MSP ID is required", fmt.Errorf("missing mspID"))
	}

	// Ensure the session directory exists
	if err := createDirectory(repository.SessionDir); err != nil {
		return handleErrorResponse(c, fiber.StatusInternalServerError, "Failed to create session directory", err)
	}

	// Create a unique session directory
	sessionID := uuid.New().String()
	sessionPath := filepath.Join(repository.SessionDir, sessionID)
	if err := createDirectory(sessionPath); err != nil {
		return handleErrorResponse(c, fiber.StatusInternalServerError, "Failed to create session directory", err)
	}

	// Save uploaded files
	files := map[string]string{
		"cert":    filepath.Join(sessionPath, "cert.pem"),
		"key":     filepath.Join(sessionPath, "key.pem"),
		"tlsCert": filepath.Join(sessionPath, "tlsCert.pem"),
	}

	for fieldName, destPath := range files {
		if err := saveFile(c, fieldName, destPath); err != nil {
			return handleErrorResponse(c, fiber.StatusInternalServerError, fmt.Sprintf("Failed to process %s", fieldName), err)
		}
	}

	// Success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Login successful",
		"data": map[string]string{
			"sessionID": sessionID,
			"mspID":     mspID,
		},
	})
}

// Logout handles user logout and session deletion.
func Logout(c *fiber.Ctx) error {
	// Get session ID from header
	sessionID := c.Get("session-id")
	if sessionID == "" {
		return handleErrorResponse(c, fiber.StatusBadRequest, "Session ID is required", fmt.Errorf("missing sessionID"))
	}

	// Delete session directory
	sessionPath := filepath.Join(repository.SessionDir, sessionID)
	if err := os.RemoveAll(sessionPath); err != nil {
		return handleErrorResponse(c, fiber.StatusInternalServerError, "Failed to delete session directory", err)
	}

	// Success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Logout successful",
		"data":    nil,
	})
}

// saveFile saves an uploaded file to the specified directory.
func saveFile(c *fiber.Ctx, fieldName, destPath string) error {
	file, err := c.FormFile(fieldName)
	if err != nil {
		return fmt.Errorf("failed to get %s file: %w", fieldName, err)
	}
	if err := c.SaveFile(file, destPath); err != nil {
		return fmt.Errorf("failed to save %s file: %w", fieldName, err)
	}
	return nil
}

// createDirectory ensures a directory exists or creates it.
func createDirectory(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}
	return nil
}

// handleErrorResponse is a reusable function for sending error responses.
func handleErrorResponse(c *fiber.Ctx, status int, message string, err error) error {
	return c.Status(status).JSON(fiber.Map{
		"success": false,
		"message": message,
		"error":   err.Error(),
		"data":    nil,
	})
}
