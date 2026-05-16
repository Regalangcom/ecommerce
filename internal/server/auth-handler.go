package server

import (
	"github.com/gin-gonic/gin"
	"github.com/regalangcom/go-shop-api/internal/dto"
	"github.com/regalangcom/go-shop-api/internal/services"
	"github.com/regalangcom/go-shop-api/internal/utils"
)

func (s *Server) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body", err)
		return
	}

	authService := services.NewAuthService(s.db, s.config)
	r, err := authService.Register(&req)
	if err != nil {
		utils.BadRequestResponse(c, "Failed to register user", err)
		return
	}

	utils.CreateResponse(c, "User registered successfully", r)
}

func (s *Server) Login(c *gin.Context) {
	var req dto.Login
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body", err)
		return
	}

	authService := services.NewAuthService(s.db, s.config)
	r, err := authService.Login(&req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to login user", err)
		return
	}

	utils.CreateResponse(c, "User logged in successfully", r)
}

func (s *Server) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body", err)
		return
	}

	authService := services.NewAuthService(s.db, s.config)
	r, err := authService.RefreshToken(&req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to refresh token", err)
		return
	}

	utils.CreateResponse(c, "Token refreshed successfully", r)
}

func (s *Server) Logout(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body", err)
		return
	}

	authService := services.NewAuthService(s.db, s.config)
	if err := authService.Logout(req.RefreshToken); err != nil {
		utils.InternalServerErrorResponse(c, "Failed to logout user", err)
		return
	}

	utils.CreateResponse(c, "User logged out successfully", nil)
}

func (s *Server) GetProfile(c *gin.Context) {

	userID := c.GetUint("user_id")
	userService := services.NewUserService(s.db)
	profile, err := userService.GetProfile(userID)

	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to get user profile", err)
		return
	}
	utils.CreateResponse(c, "User profile retrieved successfully", profile)
}

func (s *Server) UpdateProfile(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body", err)
		return
	}

	userService := services.NewUserService(s.db)
	profile, err := userService.UpdateProfile(userID, &req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update user profile", err)
		return
	}
	utils.CreateResponse(c, "User profile updated successfully", profile)
}
