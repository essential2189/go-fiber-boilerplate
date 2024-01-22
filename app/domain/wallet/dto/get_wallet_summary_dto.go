package dto

type GetWalletSummaryReq struct {
	AccountId uint64 `path:"accountId" validate:"required,numeric,gte=1"`
}

func (req *GetWalletSummaryReq) CustomValidate() error {
	return nil
}

type GetWalletSummaryRes struct {
	ProductId   int `json:"product_id"`
	ProductType int `json:"product_type"`
	Priority    int `json:"priority"`
	PassCount   int `json:"pass_count"`
}
