package repositories

import (
	"go-spp/configs"
	"go-spp/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func IsSPPAvailable(year string, month string, gradeID uuid.UUID) bool {
	var count int64
	spp := models.SPP{}
	if err := configs.DB.Model(&spp).Where("year = ? AND month = ? AND grade_id = ?", year, month, gradeID).Count(&count).Error; err != nil {
		echo.NewHTTPError(http.StatusNotFound, err.Error())
		return false
	}

	return count == 0
}
