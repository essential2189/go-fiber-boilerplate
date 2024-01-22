package exception

import (
	"errors"
	"go-boilerplate/app/core/helper/sentry"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	var customError *Error
	ok := errors.As(err, &customError)
	if ok {
		code, _ = strconv.Atoi(strings.Split(string(customError.Code), ".")[0])
		message = ErrorMessage()(customError.Code)
	}

	if r := recover(); r != nil {
		err := r.(error)
		// CHECK api error type sentry capture
		sentry.CaptureException(c, err)
		return c.Status(code).JSON(fiber.Map{"code": customError.Code, "message": message, "data": customError.Data})
	}

	return c.Status(code).JSON(fiber.Map{"error": code, "message": message, "data": customError.Data})
}
