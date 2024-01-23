package exception

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	const defaultErrorCode = fiber.StatusInternalServerError
	const defaultErrorMessage = "Internal Server Error"

	extractStatusCode := func(customError *Error) int {
		parts := strings.Split(string(customError.Code), ".")
		if len(parts) > 0 {
			if code, err := strconv.Atoi(parts[0]); err == nil {
				return code
			}
		}
		return defaultErrorCode
	}

	var customError *Error
	if errors.As(err, &customError) {
		errorResponse := GenerateErrorResponse(*customError)
		return c.Status(extractStatusCode(customError)).JSON(errorResponse)
	}

	return c.Status(defaultErrorCode).JSON(fiber.Map{
		"code":  defaultErrorCode,
		"error": defaultErrorMessage,
		"data":  nil,
	})
}
