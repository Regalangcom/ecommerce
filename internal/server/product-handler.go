package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/regalangcom/go-shop-api/internal/dto"
	"github.com/regalangcom/go-shop-api/internal/services"
	"github.com/regalangcom/go-shop-api/internal/utils"
)

// ! Category handler
func (s *Server) createCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request data", err)
		return
	}

	productService := services.NewProductService(s.db)
	category, err := productService.CreateCategory(&req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to create category", err)
		return
	}
	utils.CreateResponse(c, "Category created successfully", category)
}

func (s *Server) getCategories(c *gin.Context) {
	productService := services.NewProductService(s.db)
	categories, err := productService.GetCategory()
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to get categories", err)
		return
	}
	utils.SuccessResponse(c, "Categories retrieved successfully", categories)
}

func (s *Server) updateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "invalid category ID", err)
		return
	}

	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request data", err)
		return
	}

	productService := services.NewProductService(s.db)
	category, err := productService.UpdateCategory(uint(id), &req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to update category", err)
		return
	}
	utils.SuccessResponse(c, "Category updated successfully", category)
}

func (s *Server) deleteCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "invalid category ID", err)
		return
	}

	productService := services.NewProductService(s.db)
	if err := productService.DeleteCategory(uint(id)); err != nil {
		utils.InternalServerErrorResponse(c, "failed to delete category", err)
		return
	}
	utils.SuccessResponse(c, "Category deleted successfully", nil)
}

// ! Product handler

func (s *Server) createProduct(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request data", err)
		return
	}

	productService := services.NewProductService(s.db)
	product, err := productService.CreateProduct(&req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to create product", err)
		return
	}
	utils.CreateResponse(c, "Product created successfully", product)
}

func (s *Server) getProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	productService := services.NewProductService(s.db)
	products, meta, err := productService.GetProducts(page, limit)
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to get products", err)
		return
	}
	utils.PaginatedSuccessResponse(c, "Products retrieved successfully", products, *meta)
}

func (s *Server) getProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "invalid product ID", err)
		return
	}

	productService := services.NewProductService(s.db)
	product, err := productService.GetProduct(uint(id))
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to get product", err)
		return
	}
	utils.SuccessResponse(c, "Product retrieved successfully", product)
}

func (s *Server) updateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "invalid product ID", err)
		return
	}

	var req dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request data", err)
		return
	}

	productService := services.NewProductService(s.db)
	product, err := productService.UpdateProduct(uint(id), &req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to update product", err)
		return
	}
	utils.SuccessResponse(c, "Product updated successfully", product)
}

func (s *Server) deleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "invalid product ID", err)
		return
	}

	productService := services.NewProductService(s.db)
	if err := productService.DeleteProduct(uint(id)); err != nil {
		utils.InternalServerErrorResponse(c, "failed to delete product", err)
		return
	}
	utils.SuccessResponse(c, "Product deleted successfully", nil)
}
