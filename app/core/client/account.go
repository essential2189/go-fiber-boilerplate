package client

import (
	"encoding/json"
	"github.com/cockroachdb/errors"
	"go-boilerplate/app/core/client/dto"
	"go-boilerplate/app/core/helper"
	"go-boilerplate/app/core/helper/resty"
	"go-boilerplate/config"
)

type AccountClient interface {
	DecryptCredential(credential string) (*dto.User, error)
}

type accountClient struct {
	helper helper.Helper
}

func NewAccountClient(helper helper.Helper) AccountClient {
	return &accountClient{
		helper: helper,
	}
}

func (c accountClient) DecryptCredential(credential string) (*dto.User, error) {
	uri := config.Apis + "/account/v1/me"

	requestInfo := resty.RequestInfo{
		Uri:          uri,
		Method:       resty.MethodGET,
		Headers:      map[string]string{"Content-Type": "application/json", "credential": credential},
		Query:        map[string]string{},
		Body:         nil,
		Timeout:      0,
		RetryCount:   0,
		RetryBackOff: 0,
		IsSkipSSL:    false,
	}

	resp, err := c.helper.Resty.Request(requestInfo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get "+uri)
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("Error: " + string(resp.Body))
	}

	var resBody dto.User
	if err = json.Unmarshal(resp.Body, &resBody); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal response body")
	}

	return &resBody, nil
}
