package helpers

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
)

func init() {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found")
	}
}

func UploadToCloudinary(file *multipart.FileHeader) (string, error) {
	// Load Cloudinary configuration
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUD_NAME"),
		os.Getenv("CLOUD_KEY"),
		os.Getenv("CLOUD_SECRET"),
	)
	if err != nil {
		return "", fmt.Errorf("failed to initialize Cloudinary: %w", err)
	}

	// Open the file
	f, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	// Upload to Cloudinary
	uploadParams := uploader.UploadParams{
		Folder: "todos",
	}
	uploadResult, err := cld.Upload.Upload(context.Background(), f, uploadParams)
	if err != nil {
		return "", fmt.Errorf("upload to Cloudinary failed: %w", err)
	}

	return uploadResult.SecureURL, nil
}
