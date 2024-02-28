package errcode

const (
	Forbidden                ErrorCode = "403.000"
	AccessWithoutPermission  ErrorCode = "403.001"
	DuplicateProductPurchase ErrorCode = "403.002"
	ProductExpired           ErrorCode = "403.003"
	PartnerPassHolder        ErrorCode = "403.004"
	InAppPurchase            ErrorCode = "403.005"
)
