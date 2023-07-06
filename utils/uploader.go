package utils

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"go-image-uploader/config"
	"os"
)

func UploadToCloudinary(file *os.File, filePath string) (string, error) {
	ctx := context.Background()
	cld, err := config.SetupCloudinary()
	if err != nil {
		return "", err
	}

	uploadParams := uploader.UploadParams{
		PublicID: filePath,
	}

	result, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", err
	}

	imageUrl := result.SecureURL
	return imageUrl, nil
}
