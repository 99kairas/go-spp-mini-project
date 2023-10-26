package repositories

import (
	"go-spp/configs"
	"go-spp/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func IsPaymentAvailable(sppID uuid.UUID) bool {
	var count int64
	payment := models.Payment{}
	if err := configs.DB.Model(&payment).Where("spp_id = ?", sppID).Count(&count).Error; err != nil {
		echo.NewHTTPError(http.StatusNotFound, err.Error())
		return false
	}

	return count == 0
}

func GetSPPByID(id uuid.UUID) (models.SPP, error) {
	var spp models.SPP
	if err := configs.DB.Model(spp).Where("id = ?", id).First(&spp).Error; err != nil {
		return spp, err
	}
	return spp, nil
}

func GetAllPayments(paymentParam *models.Payment) (payment []models.Payment, err error) {
	db := configs.DB

	if err := db.Order("id").Preload("SPP.spp_id").Find(&payment).Error; err != nil {
		return nil, err
	}

	return payment, nil
}
