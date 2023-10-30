package exception

import (
	"go-boilerplate/app/core/helper/logger"
	"go-boilerplate/app/core/helper/sentry"

	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var log = logger.Get()

func ErrorHandler(c *fiber.Ctx, originErr error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	resultErr := originErr
	if errors.Is(originErr, gorm.ErrRecordNotFound) {
		resultErr = fiber.ErrNotFound
	}

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(resultErr, &e) {
		code = e.Code
	}

	if code >= 400 {
		log.Errorf("%+v", originErr)
	} else {
		log.Infof("%+v", originErr)
	}

	if r := recover(); r != nil {
		err := r.(error)
		// CHECK api error type sentry capture
		sentry.CaptureException(c, err)
		return c.Status(code).JSON(map[string]string{"error": err.Error()})
	}

	return c.Status(code).JSON(map[string]string{"error": resultErr.Error()})
}
