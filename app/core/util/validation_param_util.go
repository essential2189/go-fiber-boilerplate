package util

import (
	"go-boilerplate/app/core/helper"

	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

type Request interface {
	CustomValidate(ctx *fiber.Ctx) error
}

func GetParams(c *fiber.Ctx, param interface{}) error {
	if err := c.QueryParser(param); err != nil {
		return errors.Wrap(err, "invalid query")
	}

	err := c.BodyParser(param)
	if err != nil {
		fiberError, ok := err.(*fiber.Error)
		if !ok || fiberError != fiber.ErrUnprocessableEntity {
			return errors.Wrap(err, "invalid body")
		}
	}

	return nil
}

func CheckParameter(ctx *fiber.Ctx, helper helper.Helper, req Request) error {
	err := GetParams(ctx, req)
	if err != nil {
		return errors.WithStack(err)
	}

	err = req.CustomValidate(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	err = helper.Validator.Struct(req)

	return nil
}
