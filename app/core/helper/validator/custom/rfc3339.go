package custom

import (
	"github.com/go-playground/validator/v10"
	"time"
)

var (
	ValidationMapper = map[string]func(level validator.FieldLevel) bool{
		"rfc3339": rfc3339,
	}
)

func rfc3339(fl validator.FieldLevel) bool {
	_, err := time.Parse(time.RFC3339, fl.Field().String())
	return err == nil
}
