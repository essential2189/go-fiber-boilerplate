package validator

import (
	"github.com/cockroachdb/errors"
	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-multierror"
)

type Checker struct {
	*validator.Validate
}

func New() *Checker {
	return &Checker{validator.New()}
}

func (v *Checker) Struct(data interface{}) error {
	var result error
	errs := v.Validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			result = multierror.Append(result, err)
		}
	}
	return errors.WithStack(result)
}
