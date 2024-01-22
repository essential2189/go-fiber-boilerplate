package controller

import (
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/app"
	"go-boilerplate/app/core"
	"go-boilerplate/app/core/helper"
	"go-boilerplate/app/domain/wallet/dto"
	"go-boilerplate/app/domain/wallet/service"
)

type WalletController interface {
	Table() []app.Mapping
	GetWalletList(c *fiber.Ctx) error
}

type walletController struct {
	core   core.Modules
	helper helper.Helper
	svc    service.WalletService
}

func NewWalletController(core core.Modules, helper helper.Helper, svc service.WalletService) WalletController {
	return walletController{
		core:   core,
		helper: helper,
		svc:    svc,
	}
}

func (ctrl walletController) Table() []app.Mapping {
	return []app.Mapping{
		{Method: fiber.MethodGet, Path: "/v1/wallet/list/:accountId", Handler: ctrl.GetWalletList},
	}
}

// @Summary Get GetWalletListRes List
// @Description Get GetWalletListRes List
// @Tags GetWalletListRes
// @Accept json
// @Produce json
// @Param accountId path int true "Account ID"
// @Param pagination query dto.PaginationReq true "Pagination"
// @Success 200 {object} dto.GetWalletListRes
// @Router /v1/purchase/list/{accountId} [get]
func (ctrl walletController) GetWalletList(c *fiber.Ctx) error {
	var req dto.GetWalletListReq
	if err := ctrl.core.Base.Parameter.GetRequest(c, &req); err != nil {
		return errors.WithStack(err)
	}

	var param dto.GetWalletListParam
	if err := ctrl.core.Base.Parameter.ValidateParams(c, req, &param); err != nil {
		return errors.WithStack(err)
	}

	result, err := ctrl.svc.GetWalletList(param)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
