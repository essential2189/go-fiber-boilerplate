package validator

import (
	"github.com/cockroachdb/errors"
	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-multierror"
	"go-boilerplate/app/core/helper/validator/custom"
)

type Checker struct {
	*validator.Validate
}

func New() *Checker {
	v := &Checker{validator.New()}
	v.registerValidation()
	return v
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

func (v *Checker) registerValidation() {
	for k, f := range custom.ValidationMapper {
		err := v.RegisterValidation(k, f)
		if err != nil {
			panic(err)
		}
	}
}
