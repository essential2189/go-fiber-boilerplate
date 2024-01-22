package controller

import (
	"go-boilerplate/app"
	"go-boilerplate/app/core"
	"go-boilerplate/app/core/consts"
	"go-boilerplate/app/core/helper"
	"go-boilerplate/app/core/helper/logger"
	"go-boilerplate/app/core/helper/resty"
	"go-boilerplate/app/domain/temp/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TempController interface {
	Table() []app.Mapping
	Temp(c *fiber.Ctx) error
	Test(c *fiber.Ctx) error
}

type tempController struct {
	core    core.Modules
	helper  helper.Helper
	service service.TempService
}

func NewTempController(core core.Modules, helper helper.Helper, service service.TempService) TempController {
	return tempController{
		core:    core,
		helper:  helper,
		service: service,
	}
}

func (ctrl tempController) Table() []app.Mapping {
	return []app.Mapping{
		{fiber.MethodGet, "/temp", ctrl.Temp},
		{fiber.MethodGet, "/test", ctrl.Test},
	}
}

func (ctrl tempController) Temp(c *fiber.Ctx) error {
	requestInfo := resty.RequestInfo{
		Uri:          "http://apis/accesscheck'",
		Method:       resty.MethodGET,
		Headers:      map[string]string{"Content-Type": "application/json", consts.CredentialHeader: "asdfasdf"},
		Query:        map[string]string{},
		Body:         nil,
		Timeout:      0,
		RetryCount:   0,
		RetryBackOff: 0,
		IsSkipSSL:    false,
	}
	res, _ := ctrl.helper.Resty.Request(requestInfo)

	logger.Zap.Infof("res: " + strconv.Itoa(res.StatusCode))
	return c.SendStatus(http.StatusOK)
}

func (ctrl tempController) Test(c *fiber.Ctx) error {
	err := ctrl.service.Test()
	if err != nil {
		return err
	}
	return c.SendStatus(http.StatusOK)
}
