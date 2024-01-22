package dto

import "go-boilerplate/app/core/util"

type PaginationReq struct {
	Page    int    `query:"page" validate:"required,gte=1"`
	Limit   int    `query:"limit" validate:"required,gte=1"`
	OrderBy string `query:"order_by"`
}

type Pagination struct {
	Page    int    `query:"page" validate:"required,gte=1"`
	Limit   int    `query:"limit" validate:"required,gte=1"`
	OrderBy string `query:"order_by"`
	Offset  int
}

type PaginationResponse struct {
	CurrentPage int64       `json:"current_page"`
	TotalPage   int64       `json:"total_page"`
	TotalCount  int64       `json:"total_count"`
	ListCount   int64       `json:"list_count"`
	Data        interface{} `json:"data"`
}

func SetDefaultPagination(req *PaginationReq) Pagination {
	pagination := Pagination{
		Page:    1,
		Limit:   50,
		OrderBy: "create_ts DESC",
	}

	// Override defaults if values are provided
	if req.Page != 0 {
		pagination.Page = req.Page
	}
	if req.Limit != 0 {
		pagination.Limit = req.Limit
	}
	if req.OrderBy != "" {
		pagination.OrderBy = req.OrderBy
	}

	pagination.Offset = util.GetOffset(pagination.Page, pagination.Limit)

	return pagination
}
