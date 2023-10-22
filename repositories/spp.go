package repositories

import (
	"go-spp/configs"
	"go-spp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func IsSPPAvailable(year string, month string) bool {
	var count int64
	spp := models.SPP{}
	if err := configs.DB.Model(&spp).Where("year = ? AND month = ?", year, month).Count(&count).Error; err != nil {
		echo.NewHTTPError(http.StatusNotFound, err.Error())
		return false
	}

	return count == 0
}
