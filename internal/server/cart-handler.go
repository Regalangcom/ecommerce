package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/regalangcom/go-shop-api/internal/dto"
	"github.com/regalangcom/go-shop-api/internal/utils"
)

func (s *Server) GetCart(c *gin.Context) {
	userID := c.GetUint("user_id")

	cart, err := s.cartService.GetCart(userID)
	if err != nil {
		utils.NotFoundResponse(c, "cart not found")
		return
	}

	utils.SuccessResponse(c, "cart retrieved successfully", cart)
}

func (s *Server) AddToCart(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req dto.AddToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request data", err)
		return
	}

	cart, err := s.cartService.AddItemsToCart(userID, &req)
	if err != nil {
		utils.NotFoundResponse(c, "failed to add item to cart")
		return
	}

	utils.SuccessResponse(c, "items added  successfully", cart)
}

func (s *Server) UpdateToCart(c *gin.Context) {

	userID := c.GetUint("user_id")

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		utils.BadRequestResponse(c, "invalid cart item id", err)
		return
	}

	var req dto.UpdateCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request data", err)
		return
	}

	cart, err := s.cartService.UpdateCartItem(userID, uint(id), &req)
	if err != nil {
		utils.NotFoundResponse(c, "failed to update item to cart")
		return
	}

	utils.SuccessResponse(c, "cart items updated successfully", cart)
}

func (s *Server) RemoveToCart(c *gin.Context) {

	userID := c.GetUint("user_id")

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		utils.BadRequestResponse(c, "invalid cart item id", err)
		return
	}

	if err := s.cartService.RemoveCartItem(userID, uint(id)); err != nil {
		utils.InternalServerErrorResponse(c, "failed to delete item to cart", err)
		return
	}

	utils.SuccessResponse(c, "cart items deleted successfully", nil)
}
