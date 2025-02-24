package models

type Product struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	CategoryID  int    `json:"category_id"`
	Description string `json:"description,omitempty"`
}

type GetListProductResp struct {
	Product []Product `json:"product"`
}

type GetListProductRequest struct {
	CategoryID string `json:"category_id" binding:"required,number"`
}

type GetProductByIDRequest struct {
	ID string `json:"id" binding:"required,number"`
}
