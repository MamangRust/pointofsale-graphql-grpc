package upload_image

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/logger"
	"go.uber.org/zap"
)

type ImageUploads interface {
	EnsureUploadDirectory(uploadDir string) error
	ProcessImageUpload(file *multipart.FileHeader) (string, error)
	CleanupImageOnFailure(imagePath string)
	SaveUploadedFile(file *multipart.FileHeader, dst string) error
}

type ImageUpload struct {
	logger logger.LoggerInterface
}

func NewImageUpload(logger logger.LoggerInterface) ImageUploads {
	return &ImageUpload{logger: logger}
}

func (h *ImageUpload) EnsureUploadDirectory(uploadDir string) error {
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			h.logger.Error("failed to create upload directory",
				zap.String("directory", uploadDir),
				zap.Error(err),
			)
			return err
		}
	}
	return nil
}

func (h *ImageUpload) ProcessImageUpload(file *multipart.FileHeader) (string, error) {
	allowedTypes := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedTypes[ext] {
		return "", fmt.Errorf("invalid image type: only JPG, JPEG, PNG are allowed")
	}

	if file.Size > 5<<20 {
		return "", fmt.Errorf("image size must be less than 5MB")
	}

	uploadDir := "uploads/products"
	if err := h.EnsureUploadDirectory(uploadDir); err != nil {
		return "", fmt.Errorf("failed to prepare upload directory: %w", err)
	}

	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	imagePath := filepath.Join(uploadDir, filename)

	if err := h.SaveUploadedFile(file, imagePath); err != nil {
		h.logger.Error("failed to save uploaded file",
			zap.String("path", imagePath),
			zap.Error(err),
		)
		return "", fmt.Errorf("failed to save image: %w", err)
	}

	h.logger.Debug("uploaded image successfully",
		zap.String("path", imagePath),
		zap.Int64("size", file.Size),
	)

	return imagePath, nil
}

func (h *ImageUpload) CleanupImageOnFailure(imagePath string) {
	if removeErr := os.Remove(imagePath); removeErr != nil {
		h.logger.Debug("failed to cleanup uploaded file",
			zap.String("path", imagePath),
			zap.Error(removeErr),
		)
	}
}

func (h *ImageUpload) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer out.Close()

	if _, err = io.Copy(out, src); err != nil {
		return fmt.Errorf("failed to copy file contents: %w", err)
	}

	return nil
}
