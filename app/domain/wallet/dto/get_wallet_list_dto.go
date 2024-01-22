package dto

import (
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

type GetWalletListReq struct {
	PaginationReq
	AccountId uint64 `params:"accountId"`
}

type GetWalletListParam struct {
	Pagination
	AccountId uint64 `validate:"required,numeric,gte=1"`
}

func (param *GetWalletListParam) GenerateParam(ctx *fiber.Ctx, req interface{}) error {
	request, ok := req.(GetWalletListReq)
	if !ok {
		return errors.New("invalid request type")
	}

	param.AccountId = request.AccountId

	param.Pagination = SetDefaultPagination(&request.PaginationReq)

	return nil
}

type GetWalletListRes struct {
	// tb_wallet_product
	WalletType    int    `json:"wallet_type"`
	Priority      int    `json:"priority"`
	Renew         int    `json:"renew"`
	StartTs       string `json:"start_ts"`
	EndTs         string `json:"end_ts"`
	DownloadCount int    `json:"download_count"`
	WalletStatus  int    `json:"wallet_status"`

	// tb_bill
	CurrencyType int     `json:"currency_type"`
	BillType     int     `json:"type"`
	BillStatus   int     `json:"bill_status"`
	DeviceAgent  string  `json:"device_agent"`
	TotalAmount  float64 `json:"total_amount"`
}
