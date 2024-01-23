package exception

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ErrorHandler handles errors for the Fiber application.
func ErrorHandler(c *fiber.Ctx, err error) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	errorCode, errorResponse := determineErrorResponse(err)
	return c.Status(errorCode).JSON(errorResponse)
}

// determineErrorResponse determines the appropriate HTTP status code and error response based on the error type.
func determineErrorResponse(err error) (int, interface{}) {
	var customError *Error
	if errors.As(err, &customError) {
		errorCode := extractStatusCode(customError, fiber.StatusInternalServerError)
		return errorCode, GenerateErrorResponse(*customError)
	}

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		return fiberErr.Code, fiber.Map{
			"code":  fiberErr.Code,
			"error": fiberErr.Message,
			"data": fiber.Map{
				"error": fiberErr.Error(),
			},
		}
	}

	return fiber.StatusInternalServerError, fiber.Map{
		"code":  fiber.StatusInternalServerError,
		"error": "Internal Server Error",
		"data": fiber.Map{
			"error": err.Error(),
		},
	}
}

// extractStatusCode extracts the status code from a custom error.
func extractStatusCode(customError *Error, defaultErrorCode int) int {
	if parts := strings.Split(string(customError.Code), "."); len(parts) > 0 {
		if code, err := strconv.Atoi(parts[0]); err == nil {
			return code
		}
	}
	return defaultErrorCode
}
