package repositories

import (
	"go-spp/configs"
	"go-spp/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func IsPaymentAvailableByStudentID(sppID uuid.UUID) bool {
	var count int64
	payment := models.Payment{}
	if err := configs.DB.Model(&payment).Where("spp_id = ?", sppID).Count(&count).Error; err != nil {
		echo.NewHTTPError(http.StatusNotFound, err.Error())
		return false
	}

	return count == 0
}

func CreatePaymentByStudentID(payment *models.Payment) error {
	if err := configs.DB.Create(payment).Error; err != nil {
		return err
	}

	return nil
}

func GetSPPByID(id uuid.UUID) (models.SPP, error) {
	var spp models.SPP
	if err := configs.DB.Model(spp).Where("id = ?", id).First(&spp).Error; err != nil {
		return spp, err
	}
	return spp, nil
}

func GetGradeByID(id uuid.UUID) (models.Grade, error) {
	var grade models.Grade
	if err := configs.DB.Model(grade).Where("id = ?", id).First(&grade).Error; err != nil {
		return grade, err
	}
	return grade, nil
}

func GetAllPayments() (payment []models.Payment, err error) {
	db := configs.DB

	if err := db.Preload("Spp").Preload("Student").Preload("Admin").Find(&payment).Error; err != nil {
		return nil, err
	}

	return payment, nil
}

func IsPaymentAvailable(studentID uuid.UUID, sppID uuid.UUID) (bool, error) {
	var payment models.Payment
	if err := configs.DB.Where("student_id = ? AND spp_id = ?", studentID, sppID).First(&payment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Pembayaran belum ada
			return false, nil
		}
		return false, err
	}
	// Pembayaran sudah ada
	return true, nil
}

func IsStudentAvailable(gradeID uuid.UUID) ([]models.Student, error) {
	var students []models.Student
	if err := configs.DB.Where("grade_id = ?", gradeID).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func CreatePaymentAllStudent(payment *models.Payment) error {
	if err := configs.DB.Create(payment).Error; err != nil {
		return err
	}
	return nil
}
