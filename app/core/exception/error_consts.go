package exception

type errorCode string

const (
	ErrBadRequest errorCode = "400.001"
	// TODO: add more error codes
)

func ErrorMessage() func(errorCode) string {
	errorMessages := map[errorCode]string{
		ErrBadRequest: "Bad Request",
		// TODO: add more error messages
	}

	return func(key errorCode) string {
		return errorMessages[key]
	}
}
