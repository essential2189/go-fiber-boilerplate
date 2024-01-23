package exception

import (
	"fmt"
)

type Error struct {
	Code errorCode
	Err  error
	Data interface{}
}

type ErrorResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %s, message: %s, error: %s, data: %+v", e.Code, ErrorMessage()(e.Code), e.Err.Error(), e.Data)
}

func Wrap(code errorCode, err error, data interface{}) *Error {
	if err == nil {
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
