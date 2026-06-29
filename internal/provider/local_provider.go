package provider

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
)

type LocalUploadProvider struct {
	basePath string
}

func NewLocalProvider(basePath string) *LocalUploadProvider {
	return &LocalUploadProvider{basePath: basePath}
}

func (p *LocalUploadProvider) UploadFile(file *multipart.FileHeader, path string) (string, error) {
	fullPath := filepath.Join(p.basePath, path)

	if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
		return "", err
	}

	// Open source
	srcFile, err := file.Open()
	if err != nil {
		return "", err
	}
	defer func() {
		if cerr := srcFile.Close(); cerr != nil {
			fmt.Println("failed to close source file:", cerr)
		}
	}()

	// Create Destination
	dst, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer func() {
		if cerr := dst.Close(); cerr != nil {
			fmt.Println("failed to close destination file:", cerr)
		}
	}()

	// read from source and write to destination
	if _, err := dst.ReadFrom(srcFile); err != nil {
		return "", err
	}

	return fmt.Sprintf("/uploads/%s", path), nil
}

func (p *LocalUploadProvider) DeleteFile(path string) error {
	fullPath := filepath.Join(p.basePath, path)
	return os.Remove(fullPath)
}
