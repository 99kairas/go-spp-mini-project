package utils

import (
	"context"
	"go-spp/models/payloads"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

func Credentials() *cloudinary.Cloudinary {
	cld, _ := cloudinary.New()
	cld.Config.URL.Secure = true
	return cld
}

func UploadImageCloudBase64(req *payloads.UploadImageCloudinaryBase64Request) (imageUrl string, err error) {
	cld := Credentials()
	uuid := uuid.New()

	resp, err := cld.Upload.Upload(context.Background(), req.Image, uploader.UploadParams{
		PublicID:       "Go-SPP/" + uuid.String(),
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true),
	})

	if err != nil {
		return "", err
	}

	imageUrl = resp.SecureURL

	return imageUrl, nil
}
