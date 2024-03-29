package repositories

import (
	"go-spp/configs"
	"go-spp/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func IsUsernameAvailable(username string) bool {
	var count int64
	admin := models.Admin{}
	if err := configs.DB.Model(&admin).Where("username = ?", username).Count(&count).Error; err != nil {
		echo.NewHTTPError(http.StatusNotFound, err.Error())
		return false
	}

	return count == 0
}

func CreateAdmin(admin *models.Admin) error {
	if err := configs.DB.Create(admin).Error; err != nil {
		return err
	}

	return nil
}

func GetAdmin(username string) (admin *models.Admin, err error) {
	if err := configs.DB.Where("username = ?", username).First(&admin).Error; err != nil {
		return admin, err
	}
	return admin, nil
}

func CreateSPP(spp *models.SPP) error {
	if err := configs.DB.Create(spp).Error; err != nil {
		return err
	}

	return nil
}

func GetStudentByID(id uuid.UUID) (student *models.Student, err error) {
	if err := configs.DB.Where("id = ?", id).Preload("Grade").First(&student).Error; err != nil {
		return student, err
	}

	return student, nil
}

func GetAllStudent() (student []models.Student, err error) {
	if err := configs.DB.Preload("Grade").Find(&student).Error; err != nil {
		return nil, err
	}

	return student, nil
}

func DeleteStudent(student *models.Student) error {
	if err := configs.DB.Delete(&student).Error; err != nil {
		return err
	}

	return nil
}
