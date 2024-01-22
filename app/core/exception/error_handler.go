package exception

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"
	var data interface{}

	var customError *Error
	ok := errors.As(err, &customError)
	if ok {
		code, _ = strconv.Atoi(strings.Split(string(customError.Code), ".")[0])
		message = ErrorMessage()(customError.Code)
		data = customError.Data
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(code).JSON(fiber.Map{"error": code, "message": message, "data": data})
}
