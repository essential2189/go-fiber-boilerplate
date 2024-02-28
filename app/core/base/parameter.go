package base

import (
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/app/core/exception"
	"go-boilerplate/app/core/exception/errcode"
	"go-boilerplate/app/core/helper"
)

type Parameter interface {
	GetRequest(ctx *fiber.Ctx, param interface{}) error
	ValidateParams(ctx *fiber.Ctx, req interface{}, param param) error
}

type getParameter struct {
	helper helper.Helper
}

func NewGetParameter(helper helper.Helper) Parameter {
	return getParameter{
		helper: helper,
	}
}

func (gp getParameter) GetRequest(ctx *fiber.Ctx, param interface{}) error {
	if err := ctx.ParamsParser(param); err != nil {
		return exception.WithData(errcode.InvalidParameter, err, fiber.Map{"error": err.Error()})
	}

	if err := ctx.ReqHeaderParser(param); err != nil {
		return exception.WithData(errcode.InvalidParameter, err, fiber.Map{"error": err.Error()})
	}

	if err := ctx.QueryParser(param); err != nil {
		return exception.WithData(errcode.InvalidParameter, err, fiber.Map{"error": err.Error()})
	}

	if ctx.GetReqHeaders()["Content-Type"][0] == "application/json" {
		err := ctx.BodyParser(param)
		if err != nil {
			return exception.WithData(errcode.InvalidParameter, err, fiber.Map{"error": err.Error()})
		}
	}

	return nil
}

func (gp getParameter) ValidateParams(ctx *fiber.Ctx, req interface{}, param param) error {
	err := param.GenerateParam(ctx, req)
	if err != nil {
		return errors.WithStack(err)
	}

	err = gp.helper.Validator.Struct(param)

	return errors.WithStack(err)
}

type param interface {
	GenerateParam(ctx *fiber.Ctx, req interface{}) error
}
