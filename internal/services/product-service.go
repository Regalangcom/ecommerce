package services

import (
	"github.com/regalangcom/go-shop-api/internal/dto"
	"github.com/regalangcom/go-shop-api/internal/models"
	"github.com/regalangcom/go-shop-api/internal/utils"
	"gorm.io/gorm"
)

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{db: db}
}

// !!! NOTE: When creating a category, we should ensure that the name is unique to prevent duplicate categories. This can be done by adding a unique constraint on the Name field in the database and handling the error appropriately when trying to create a category with a name that already exists. Additionally, we should validate that the name is not empty and meets any other business requirements before attempting to create the category.
func (s *ProductService) CreateCategory(req *dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	category := models.Category{
		Name:        req.Name,
		Description: req.Description,
	}
	if err := s.db.Create(&category).Error; err != nil {
		return nil, err
	}
	return &dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (s *ProductService) GetCategory() ([]dto.CategoryResponse, error) {
	var categories []models.Category

	if err := s.db.Where("is_active = ?", true).Find(&categories).Error; err != nil {
		return nil, err
	}
	response := make([]dto.CategoryResponse, len(categories))
	for i := range categories {
		response[i] = dto.CategoryResponse{

			ID:          categories[i].ID,
			Name:        categories[i].Name,
			Description: categories[i].Description,
			IsActive:    categories[i].IsActive,
		}
	}
	return response, nil
}

func (s *ProductService) UpdateCategory(id uint, req *dto.UpdateCategoryRequest) (*dto.CategoryResponse, error) {
	var category models.Category
	if err := s.db.First(&category, id).Error; err != nil {
		return nil, err
	}

	category.Name = req.Name
	category.Description = req.Description
	if req.IsActive != nil {
		category.IsActive = *req.IsActive
	}

	if err := s.db.Save(&category).Error; err != nil {
		return nil, err
	}

	return &dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		IsActive:    category.IsActive,
	}, nil
}

func (s *ProductService) DeleteCategory(id uint) error {
	return s.db.Delete(&models.Category{}, id).Error
}

// !!! NOTE: When creating a product, we should also validate that the provided CategoryID exists and is active. This can be done by querying the database for the category before creating the product. If the category does not exist or is not active, we should return an appropriate error message. This validation step is crucial to maintain data integrity and ensure that products are associated with valid categories.
func (s *ProductService) CreateProduct(req *dto.CreateProductRequest) (*dto.ProductResponse, error) {
	product := models.Product{
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		SKU:         req.SKU,
	}
	if err := s.db.Create(&product).Error; err != nil {
		return nil, err
	}
	return s.GetProduct(product.ID)
}

func (s *ProductService) GetProducts(page, limit int) ([]dto.ProductResponse, *utils.PaginatedMeta, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offest := (page - 1) * limit
	var products []models.Product
	var total int64

	s.db.Model(&models.Product{}).Where("is_active = ?", true).Count(&total)
	if err := s.db.Preload("Category").Preload("Images").
		Where("is_active = ?", true).
		Offset(offest).
		Limit(limit).
		Find(&products).Error; err != nil {
		return nil, nil, err
	}

	response := make([]dto.ProductResponse, len(products))
	for i := range products {
		response[i] = s.ConvertToProductResponse(&products[i])
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))
	meta := &utils.PaginatedMeta{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}

	return response, meta, nil
}

func (s *ProductService) GetProduct(id uint) (*dto.ProductResponse, error) {
	var product models.Product
	if err := s.db.Preload("Category").
		Preload("Images").
		First(&product, id).Error; err != nil {
		return nil, err
	}
	response := s.ConvertToProductResponse(&product)
	return &response, nil
}

func (s *ProductService) UpdateProduct(id uint, req *dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	var product models.Product
	if err := s.db.First(&product, id).Error; err != nil {
		return nil, err
	}

	product.CategoryID = req.CategoryID
	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = req.Stock
	if req.IsActive != nil {
		product.IsActive = *req.IsActive
	}

	if err := s.db.Save(&product).Error; err != nil {
		return nil, err
	}

	return s.GetProduct(id)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.db.Delete(&models.Product{}, id).Error
}

func (s *ProductService) ConvertToProductResponse(product *models.Product) dto.ProductResponse {
	images := make([]dto.ProductImageResponse, len(product.Images))
	for i := range product.Images {
		images[i] = dto.ProductImageResponse{
			ID:        product.Images[i].ID,
			URL:       product.Images[i].URL,
			IsPrimary: product.Images[i].IsPrimary,
			AltText:   product.Images[i].AltText,
		}
	}
	return dto.ProductResponse{
		ID:          product.ID,
		CategoryID:  product.CategoryID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		SKU:         product.SKU,
		IsActive:    product.IsActive,
		Category: dto.CategoryResponse{
			ID:          product.Category.ID,
			Name:        product.Category.Name,
			Description: product.Category.Description,
			IsActive:    product.Category.IsActive,
		},
		Images: images,
	}
}
