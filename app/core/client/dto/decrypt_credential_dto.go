package dto

import (
	"time"
)

type User struct {
	IdList    []IdInfo   `json:"id_list"`
	Gender    string     `json:"gender"`
	JoinDate  time.Time  `json:"join_date"`
	Birth     time.Time  `json:"birth"`
	Personal  Personal   `json:"personal"`
	Push      Push       `json:"push"`
	AccountId uint64     `json:"account_id"`
	Phone     string     `json:"phone"`
	Terms     []TermInfo `json:"terms"`
	ProfileId string     `json:"profile_id"`
	Uno       int        `json:"uno"`
	Name      string     `json:"name"`
	Adult     Adult      `json:"adult"`
	Email     string     `json:"email"`
}

type IdInfo struct {
	LoginType  string `json:"login_type"`
	JoinDate   bool   `json:"join_date"`
	ProviderId int    `json:"provider_id"`
	LoginDesc  string `json:"login_desc"`
	ID         string `json:"id"`
}

type Personal struct {
	IsAuth   bool      `json:"is_auth"`
	AuthDate time.Time `json:"auth_date"`
	Type     string    `json:"type"`
}

type Push struct {
	ServiceAgree       bool      `json:"service_agree"`
	MarketingAgreeDate time.Time `json:"marketing_agree_date"`
	MarketingAgree     bool      `json:"marketing_agree"`
	ServiceAgreeDate   time.Time `json:"service_agree_date"`
}

type TermInfo struct {
	TermTypeId int       `json:"term_type_id"`
	AgreeDate  time.Time `json:"agree_date"`
	TermID     int       `json:"term_id"`
	IsAgree    bool      `json:"is_agree"`
}

type Adult struct {
	VerifyDate time.Time `json:"verify_date"`
	IsAdult    bool      `json:"is_adult"`
}
