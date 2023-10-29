package repositories

import (
	"go-spp/configs"
	"go-spp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func IsNISAvailable(nis string) bool {
	var count int64
	student := models.Student{}
	if err := configs.DB.Model(&student).Where("nis = ?", nis).Count(&count).Error; err != nil {
		echo.NewHTTPError(http.StatusNotFound, err.Error())
		return false
	}

	return count == 0
}

func CreateStudent(student *models.Student) error {
	if err := configs.DB.Create(student).Error; err != nil {
		return err
	}
	return nil
}

func GetStudent(nis string) (student *models.Student, err error) {
	if err := configs.DB.Where("nis = ?", nis).First(&student).Error; err != nil {
		return student, err
	}
	return student, nil
}

func UpdateStudent(student *models.Student) error {
	if err := configs.DB.Updates(&student).Error; err != nil {
		return err
	}

	return nil
}
