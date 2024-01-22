package dto

type WalletReq struct {
	AccountId uint64 `json:"account_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}
