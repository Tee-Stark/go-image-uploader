package utils

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"go-image-uploader/config"
	"mime/multipart"
)

func UploadToCloudinary(file multipart.File, filePath string) (string, error) {
	ctx := context.Background()
	cld, err := config.SetupCloudinary()
	if err != nil {
		return "", err
	}

	// create upload params
	uploadParams := uploader.UploadParams{
		PublicID:     filePath,
		ResourceType: "image",
	}

	result, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", err
	}

	imageUrl := result.SecureURL
	return imageUrl, nil
}
