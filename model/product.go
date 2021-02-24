package model

type ProductDto struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Quantity    uint    `json:"quantity"`
	Price       float64 `json:"price"`
}

type ProductsByPaginationResponse struct {
	PageSize   uint         `json:"pageSize"`
	PageNumber uint         `json:"pageNumber"`
	TotalCount uint         `json:"totalCount"`
	Products   []ProductDto `json:"products"`
}
