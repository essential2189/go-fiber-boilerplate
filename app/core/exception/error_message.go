package exception

import "go-boilerplate/app/core/exception/errcode"

func ErrorMessage() func(errcode.ErrorCode) string {
	errorMessages := map[errcode.ErrorCode]string{
		errcode.BadRequest: "Bad Request",
		// TODO: add more error messages
	}

	return func(key errcode.ErrorCode) string {
		return errorMessages[key]
	}
}
