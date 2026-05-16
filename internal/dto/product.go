package dto

type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

type CategoryResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

type CreateProductRequest struct {
	CategoryID  uint    `json:"category_id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	SKU         string  `json:"sku" binding:"required"`
}

type UpdateProductRequest struct {
	CategoryID  uint    `json:"category_id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	IsActive    bool    `json:"is_active"`
}

type ProductResponse struct {
	ID          uint                   `json:"id"`
	CategoryID  uint                   `json:"category_id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Price       float64                `json:"price"`
	Stock       int                    `json:"stock"`
	SKU         string                 `json:"sku"`
	IsActive    bool                   `json:"is_active"`
	Category    CategoryResponse       `json:"category"`
	Images      []ProductImageResponse `json:"images"`
}

type ProductImageResponse struct {
	ID uint `json:"id"`
	// ProductID uint   `json:"product_id"`
	URL       string `json:"url"`
	IsPrimary bool   `json:"is_primary"`
	AltText   string `json:"alt_text"`
}
