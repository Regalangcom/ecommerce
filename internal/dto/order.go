package dto

type AddToCartRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}

type UpdateCartItemRequest struct {
	Quantity int `json:"quantity" binding:"required"`
}

type CartResponse struct {
	ID        uint               `json:"id"`
	UserID    uint               `json:"user_id"`
	CartItems []CartItemResponse `json:"items"`
	Total     float64            `json:"total_price"`
}

type CartItemResponse struct {
	ID       uint    `json:"product_id"`
	Product  string  `json:"product_name"`
	Quantity int     `json:"quantity"`
	SubTotal float64 `json:"subtotal"`
}

type OrderResponse struct {
	ID          uint                `json:"id"`
	UserID      uint                `json:"user_id"`
	Status      string              `json:"status"`
	TotalAmount float64             `json:"total_amount"`
	OrderItems  []OrderItemResponse `json:"order_items"`
	CreateAt    string              `json:"created_at"`
}

type OrderItemResponse struct {
	ID       uint    `json:"product_id"`
	Product  string  `json:"product_name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
