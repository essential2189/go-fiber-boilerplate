package exception

import "fmt"

type Error struct {
	Code errorCode   `json:"code"`
	Err  error       `json:"error"`
	Data interface{} `json:"data"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %s, error: %s, data: %+v", e.Code, e.Err.Error(), e.Data)
}

func Wrap(code errorCode, err error, data interface{}) *Error {
	return &Error{
		Code: code,
		Err:  err,
		Data: data,
	}
}
