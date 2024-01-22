package dto

type ProductRes struct {
	ProductCode       string `json:"product_code"`
	ProductId         int32  `json:"product_id"`
	PaymentId         int32  `json:"payment_id"`
	CurrencyCode      string `json:"currency_code"` //통화 타입 code
	CurrencyName      string `json:"currency_name"` //통화 타입 code
	TotalAmount       int    `json:"total_amount"`
	Renew             int    `json:"renew"`
	ProductGroupType  string `json:"product_group_type"`
	ProductType       string `json:"product_type"`
	StoreType         int    `json:"store_type"`
	StoreName         string `json:"store_name"`
	DiscountType      string `json:"discount_type"`
	DiscountAmount    int    `json:"discount_amount"`
	CouponId          int    `json:"coupon_id"`
	ExternCode        string `json:"extern_code"`
	Duration          int    `json:"duration"`
	PeriodType        string `json:"period_type"`
	StreamingDayCount int    `json:"streaming_day_count"`
}
