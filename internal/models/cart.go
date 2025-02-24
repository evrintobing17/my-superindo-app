package models

type Cart struct {
	ProductID string `json:"id_product" binding:"required,number"`
	Total     string `json:"total" binding:"required,number"`
}

type AddToCardRequest struct {
	ProductID string `json:"id_product" binding:"required,number"`
	Total     string `json:"total" binding:"required,number"`
}
