package client

import (
	"encoding/json"
	"github.com/cockroachdb/errors"
	"go-boilerplate/app/core/client/dto"
	"go-boilerplate/app/core/helper"
	"go-boilerplate/app/core/helper/resty"
	"go-boilerplate/config"
	"strconv"
)

type PurchaseClient interface {
	GetProductInfoByPaymentId(paymentId int) (*dto.ProductRes, error)
}

type purchaseClient struct {
	config *config.Config
	helper helper.Helper
}

func NewPurchaseClient(config *config.Config, helper helper.Helper) PurchaseClient {
	return &purchaseClient{
		config: config,
		helper: helper,
	}
}
func (c purchaseClient) GetProductInfoByPaymentId(paymentId int) (*dto.ProductRes, error) {
	reqInfo := resty.RequestInfo{
		Uri:    config.Apis + "/products/defrayment/detail/" + strconv.Itoa(paymentId),
		Method: resty.MethodGET,
	}

	resp, err := c.helper.Resty.Request(reqInfo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get product info")
	}

	res := &dto.ProductRes{}
	if resp.StatusCode != 200 {
		return nil, errors.New("Error: " + string(resp.Body))
	}

	if err = json.Unmarshal(resp.Body, res); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal product info")
	}

	return res, nil
}
