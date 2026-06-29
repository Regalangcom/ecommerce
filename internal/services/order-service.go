package services

import (
	"errors"
	"fmt"

	"github.com/regalangcom/go-shop-api/internal/dto"
	"github.com/regalangcom/go-shop-api/internal/models"
	"gorm.io/gorm"
)

type OrderService struct {
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{db: db}
}

func (s *OrderService) CreateOrder(userID uint) (*dto.OrderResponse, error) {
	var OrderResponse *dto.OrderResponse

	// rollback
	err := s.db.Transaction(func(tx *gorm.DB) error {

		var cart models.Cart
		if err := tx.
			Preload("CartItems.Product").
			Where("user_id = ?", userID).
			First(&cart).Error; err != nil {
			return errors.New("cart not found")
		}

		if len(cart.CartItems) == 0 {
			return errors.New("cart is empty")
		}

		// =============================
		// CALCULATE TOTAL
		// =============================
		var totalAmount float64
		var orderItems []models.OrderItem

		for i := range cart.CartItems {
			cartItem := &cart.CartItems[i]

			if cartItem.Product.Stock < cartItem.Quantity {
				return fmt.Errorf("insuficient stock for product: %s", cartItem.Product.Name)
			}

			itemTotal := float64(cartItem.Quantity) * cartItem.Product.Price
			totalAmount += itemTotal

			orderItems = append(orderItems, models.OrderItem{
				ProductID: cartItem.ProductID,
				Quantity:  cartItem.Quantity,
				Price:     cartItem.Product.Price,
			})

			// update product stock
			cartItem.Product.Stock -= cartItem.Quantity
			if err := tx.Save(&cartItem.Product).Error; err != nil {
				return err
			}

			// =============================
			// CREATE ORDER
			// =============================

			order := models.Order{
				UserID:      userID,
				Status:      string(models.OrderStatusPending),
				TotalAmount: totalAmount,
				OrderItems:  orderItems,
			}

			if err := tx.Create(&order).Error; err != nil {
				return err
			}

			// =============================
			// CLEAR CART
			// =============================

			if err := tx.
				Where("cart_id = ?", cart.ID).
				Delete(&models.CartItem{}).Error; err != nil {
				return err
			}

			response, err := s.getOrderResponse(tx, order.ID)
			if err != nil {
				return err
			}

			OrderResponse = response

			return nil
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return OrderResponse, nil

}

func (s *OrderService) convertToOrderResponse(order *models.Order) dto.OrderResponse {
	orderItems := make([]dto.OrderItemResponse, len(order.OrderItems))

	for i := range order.OrderItems {
		item := order.OrderItems[i]
	}

}
