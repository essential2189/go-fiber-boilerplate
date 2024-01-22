package service

import (
	"github.com/cockroachdb/errors"
	"go-boilerplate/app/core/exception"
)

type TempService interface {
	Test() error
}

type tempService struct {
}

func NewTempService() TempService {
	return tempService{}
}

func (s tempService) Test() error {
	err := errors.New("test error")

	return exception.Wrap(exception.ErrBadRequest, errors.WithStack(err), nil)
}
