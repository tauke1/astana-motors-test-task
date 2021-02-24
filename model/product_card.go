package model

type ProductCardDto struct {
	ProductID uint       `json:"product_id"`
	Quantity  uint       `json:"quantity"`
	Product   ProductDto `json:"product"`
}

type ChangeProductCardItemQuantityRequest struct {
	Quantity  uint `json:"quantity"`
	ProductID uint `json:"product_id"`
}

type DeleteProductCardItemRequest struct {
	ProductID uint `json:"product_id"`
}
