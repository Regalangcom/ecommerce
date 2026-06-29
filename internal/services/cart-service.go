package services

import (
	"errors"

	"github.com/regalangcom/go-shop-api/internal/dto"
	"github.com/regalangcom/go-shop-api/internal/models"
	"gorm.io/gorm"
)

type CartService struct {
	db *gorm.DB
}

// constructor
func NewCartService(db *gorm.DB) *CartService {
	return &CartService{db: db}
}

func (s *CartService) GetCart(userID uint) (*dto.CartResponse, error) {
	var Cart models.Cart
	// CartItems.Product.Category = this is nested relation
	err := s.db.Preload("CartItems.Product.Category").
		Where("user_id = ?", userID).First(&Cart).Error
	if err != nil {
		return nil, err
	}

	return s.convertToCartResponses(&Cart), nil
}

func (s *CartService) AddItemsToCart(userID uint, req *dto.AddToCartRequest) (*dto.CartResponse, error) {
	var product models.Product
	if err := s.db.First(&product, req.ProductID).Error; err != nil {
		return nil, errors.New("product not found")
	}

	if product.Stock < req.Quantity {
		return nil, errors.New("insuficient stock")
	}

	// GET or Create cart
	var cart models.Cart
	if err := s.db.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		cart := models.Cart{UserID: userID}
		if err := s.db.Create(&cart).Error; err != nil {
			return nil, err
		}
	}

	// check if item already in the cart
	var cartItem models.CartItem
	if err := s.db.Where("cart_id = ? AND product_id = ?", cart.ID, req.ProductID).
		First(&cartItem).Error; err != nil {
		// Create CartItem
		cartItem = models.CartItem{
			CartID:    cart.ID,
			ProductID: req.ProductID,
			Quantity:  req.Quantity,
		}
		s.db.Create(&cartItem)
	} else {
		// update existing cart items
		cartItem.Quantity += req.Quantity
		if cartItem.Quantity > product.Stock {
			return nil, errors.New("insuficient stock")
		}
		s.db.Save(&cartItem)
	}
	return s.GetCart(userID)
}

func (s *CartService) UpdateCartItem(userID uint, itemID uint, req *dto.UpdateCartItemRequest) (*dto.CartResponse, error) {
	var cartItem models.CartItem
	if err := s.db.Joins("JOIN carts ON cart_items.cart_id = carts.id").
		Where("cart_items.id  = ? AND carts.user_id = ?", itemID, userID).
		First(&cartItem).Error; err != nil {
		return nil, errors.New("cart item not found")
	}
	var product models.Product

	if err := s.db.First(&product, cartItem.ProductID).Error; err != nil {
		return nil, errors.New("product not found")
	}
	if product.Stock < req.Quantity {
		return nil, errors.New("insuficient stock")
	}

	cartItem.Quantity = req.Quantity
	if err := s.db.Save(&cartItem).Error; err != nil {
		return nil, err
	}

	return s.GetCart(userID)

}

func (s *CartService) RemoveCartItem(userID uint, itemID uint) error {
	return s.db.Where("id = ? AND cart_id IN (?)", itemID, s.db.Select("id").
		Table("carts").
		Where("user_id = ?", userID)).
		Delete(&models.CartItem{}).Error
}

func (s *CartService) convertToCartResponses(cart *models.Cart) *dto.CartResponse {
	cartItems := make([]dto.CartItemResponse, len(cart.CartItems)) /* memory allocation */

	var total float64
	for i := range cartItems {
		subTotal := float64(cart.CartItems[i].Quantity) * cart.CartItems[i].Product.Price
		total += subTotal

		cartItems[i] = dto.CartItemResponse{
			ID: cart.CartItems[i].ID,
			Product: dto.ProductResponse{
				ID:          cart.CartItems[i].Product.ID,
				CategoryID:  cart.CartItems[i].Product.CategoryID,
				Name:        cart.CartItems[i].Product.Name,
				Description: cart.CartItems[i].Product.Description,
				Price:       cart.CartItems[i].Product.Price,
				Stock:       cart.CartItems[i].Product.Stock,
				SKU:         cart.CartItems[i].Product.SKU,
				IsActive:    cart.CartItems[i].Product.IsActive,
				Category: dto.CategoryResponse{
					ID:          cart.CartItems[i].Product.Category.ID,
					Name:        cart.CartItems[i].Product.Category.Name,
					Description: cart.CartItems[i].Product.Category.Name,
					IsActive:    cart.CartItems[i].Product.Category.IsActive,
				},
			},
			Quantity: cart.CartItems[i].Quantity,
			SubTotal: subTotal,
		}
	}
	return &dto.CartResponse{
		ID:        cart.ID,
		UserID:    cart.UserID,
		CartItems: cartItems,
		Total:     total,
	}
}
