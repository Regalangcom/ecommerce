package services

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/regalangcom/go-shop-api/internal/interfaces"
)

type UploadService struct {
	provider interfaces.UploadProvider
}

func NewUploadService(provider interfaces.UploadProvider) *UploadService {
	return &UploadService{provider: provider}
}

func (s *UploadService) UploadProductImage(productID uint, file *multipart.FileHeader) (string, error) {
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !isValidImageExt(ext) {
		return "", fmt.Errorf("invalid type: %s", ext)
	}

	newFileName := uuid.New().String() + ext

	path := fmt.Sprintf("products/%d/%s", productID, newFileName)
	return s.provider.UploadFile(file, path)
}

func isValidImageExt(ext string) bool {
	validExts := []string{".jpg", ".jpeg", ".png", ".gif"}
	for _, validExt := range validExts {
		if ext == validExt {
			return true
		}
	}
	return false
}
