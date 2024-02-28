package errcode

const (
	BadRequest                     ErrorCode = "400.000"
	MissingRequiredParameter       ErrorCode = "400.001"
	DataNotFound                   ErrorCode = "400.002"
	UserNameAndEmailNotMatch       ErrorCode = "400.003"
	InvalidPassword                ErrorCode = "400.004"
	UnusablePassword               ErrorCode = "400.005"
	AccountLockedByInvalidPassword ErrorCode = "400.006"
	InvalidParameter               ErrorCode = "400.007"
	PasswordExpired                ErrorCode = "400.008"
	UnableWithdrawMember           ErrorCode = "400.009"
	CannotModify                   ErrorCode = "400.010"
	InitialPassword                ErrorCode = "400.011"
	ExceededIPRegistrationLimit    ErrorCode = "400.012"
	DuplicateData                  ErrorCode = "400.013"
)
