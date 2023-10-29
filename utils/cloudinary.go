package utils

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

func Credentials() *cloudinary.Cloudinary {
	cld, _ := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	cld.Config.URL.Secure = true
	return cld
}

func UploadPaymentPhoto(fileHeader *multipart.FileHeader) (imageUrl string, err error) {
	file, _ := fileHeader.Open()
	uid := uuid.New()

	cld := Credentials()

	resp, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		PublicID:       "Go-SPP/" + uid.String(),
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true),
	})

	if err != nil {
		return
	}

	imageUrl = resp.SecureURL

	return imageUrl, nil
}
