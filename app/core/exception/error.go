package exception

import (
	"encoding/json"
	"go-boilerplate/app/core/exception/errcode"
	"reflect"
)

type Error struct {
	Code errcode.ErrorCode
	Err  error
	Data interface{}
}

type ErrorResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorLog struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

func (e *Error) Error() string {
	log := ErrorLog{
		Code:    string(e.Code),
		Message: ErrorMessage()(e.Code),
		Error:   e.Err.Error(),
		Data:    e.Data,
	}

	jsonBytes, _ := json.Marshal(log)

	return string(jsonBytes)
}

func Wrap(code errcode.ErrorCode, err error) error {
	if err == nil || (reflect.ValueOf(err).Kind() == reflect.Ptr && reflect.ValueOf(err).IsNil()) {
		return nil
	}

	return &Error{
		Code: code,
		Err:  err,
	}
}

func WithData(code errcode.ErrorCode, err error, data interface{}) error {
	if err == nil || (reflect.ValueOf(err).Kind() == reflect.Ptr && reflect.ValueOf(err).IsNil()) {
		return nil
	}

	return &Error{
		Code: code,
		Err:  err,
		Data: data,
	}
}

func GenerateErrorResponse(e Error) *ErrorResponse {
	return &ErrorResponse{
		Code:    string(e.Code),
		Message: ErrorMessage()(e.Code),
		Data:    e.Data,
	}
}
