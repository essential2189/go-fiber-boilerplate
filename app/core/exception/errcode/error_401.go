package errcode

const (
	Unauthorized ErrorCode = "401.000"
	InvalidToken ErrorCode = "401.001"
	ExpiredToken ErrorCode = "401.002"
	UnregisterIP ErrorCode = "401.003"
)
