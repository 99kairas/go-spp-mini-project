package usecase

import (
	"errors"
	"go-spp/models"
	"go-spp/models/payloads"
	"go-spp/repositories"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateStudent(req *payloads.CreateStudentRequest) (resp payloads.CreateStudentResponse, err error) {
	if !repositories.IsNISAvailable(req.NIS) {
		return resp, errors.New("nomor induk siswa is already registered")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return
	}

	newStudent := &models.Student{
		ID:          uuid.New(),
		NIS:         req.NIS,
		Password:    string(passwordHash),
		FirstName:   req.Firstname,
		LastName:    req.Lastname,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		GradeID:     req.GradeID,
	}

	err = repositories.CreateStudent(newStudent)

	if err != nil {
		return
	}

	resp = payloads.CreateStudentResponse{
		ID:        newStudent.ID,
		NIS:       newStudent.NIS,
		Firstname: newStudent.FirstName,
		Lastname:  newStudent.LastName,
		Address:   newStudent.Address,
		GradeID:   newStudent.GradeID,
	}

	return
}

func UpdatePassword(id uuid.UUID, req *payloads.UpdatePasswordRequest) (res payloads.UpdatePasswordRequest, err error) {
	student, err := repositories.GetStudentByID(id)
	if err != nil {
		return res, errors.New("student not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(req.CurrentPassword))
	if err != nil {
		return res, errors.New("current password is incorrect")
	}

	if req.ConfirmPassword != req.Password {
		return res, errors.New("passwords do not match")
	}

	if req.CurrentPassword == req.Password {
		return res, errors.New("new password must be different from the current password")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return res, err
	}

	student.Password = string(passwordHash)

	err = repositories.UpdateStudent(student)
	if err != nil {
		return res, errors.New("can't update password")
	}
	return res, nil
}

func UpdateProfile(student *models.Student, req *payloads.UpdateProfileStudentRequest) (res payloads.UpdateProfileStudentResponse, err error) {
	if req.BirthDate != "" {
		if student.BirthDate != nil {
			return res, errors.New("birth date can only be updated once")
		}

		birthDate, err := time.Parse("02/01/2006", req.BirthDate)
		if err != nil {
			return res, errors.New("error on birth date")
		}
		student.BirthDate = &birthDate
	}

	if req.PhoneNumber != "" {
		student.PhoneNumber = "+62 " + req.PhoneNumber
	}

	if req.Address != "" {
		student.Address = req.Address
	}

	err = repositories.UpdateStudent(student)
	if err != nil {
		return res, errors.New("can't update profile")
	}

	res = payloads.UpdateProfileStudentResponse{
		BirthDate:   student.BirthDate,
		PhoneNumber: student.PhoneNumber,
		Address:     student.Address,
	}

	return
}

func UploadImage(paymentID uuid.UUID, image string) (imageUrl string, err error) {
	imageUrl, err = repositories.UploadImage(paymentID, image)
	if err != nil {
		return "", err
	}

	return imageUrl, nil
}

func GetStudent(studentID uuid.UUID) (payment []models.Payment, err error) {
	payment, err = repositories.GetPaymentsID(studentID)
	if err != nil {
		return nil, errors.New("student not found")
	}

	return payment, nil
}

func GetPayments(studentID uuid.UUID) (res []payloads.GetAllPaymentsResponse, err error) {
	payments, err := repositories.GetPaymentsID(studentID)
	if err != nil {
		return nil, errors.New("payments not found")
	}

	for _, student := range payments {
		res = append(res, payloads.GetAllPaymentsResponse{
			ID: student.ID,
			Spp: payloads.SPPResponse{
				ID:    student.Spp.ID,
				Year:  student.Spp.Year,
				Month: student.Spp.Month,
			},
			Student: payloads.StudentResponse{
				ID:   student.Student.ID,
				NIS:  student.Student.NIS,
				Name: student.Student.FirstName + " " + student.Student.LastName,
			},
			Admin: payloads.AdminResponse{
				ID:       student.Admin.ID,
				Username: student.Admin.Username,
				Name:     student.Admin.Name,
			},
			TotalAmount:   student.TotalAmount,
			PaymentPhoto:  student.PaymentPhoto,
			PaymentStatus: student.PaymentStatus,
			CreatedAt:     &student.CreatedAt,
			UpdatedAt:     &student.UpdatedAt,
		})
	}

	return res, nil
}

func GetDetailPayments(paymentID uuid.UUID) (res []payloads.GetAllPaymentsResponse, err error) {
	payments, err := repositories.GetDetailPaymentsID(paymentID)
	if err != nil {
		return nil, errors.New("payments not found")
	}

	for _, student := range payments {
		res = append(res, payloads.GetAllPaymentsResponse{
			ID: student.ID,
			Spp: payloads.SPPResponse{
				ID:    student.Spp.ID,
				Year:  student.Spp.Year,
				Month: student.Spp.Month,
			},
			Student: payloads.StudentResponse{
				ID:   student.Student.ID,
				NIS:  student.Student.NIS,
				Name: student.Student.FirstName + " " + student.Student.LastName,
			},
			Admin: payloads.AdminResponse{
				ID:       student.Admin.ID,
				Username: student.Admin.Username,
				Name:     student.Admin.Name,
			},
			TotalAmount:   student.TotalAmount,
			PaymentPhoto:  student.PaymentPhoto,
			PaymentStatus: student.PaymentStatus,
			CreatedAt:     &student.CreatedAt,
			UpdatedAt:     &student.UpdatedAt,
		})
	}

	return res, nil
}

func DeleteStudent(studentID uuid.UUID) error {
	student, err := repositories.GetStudentByID(studentID)
	if err != nil {
		return errors.New("student not found")
	}

	if err := repositories.DeletePaymentsByStudentID(studentID); err != nil {
		return err
	}

	err = repositories.DeleteStudent(student)
	if err != nil {
		return errors.New("failed to delete the student")
	}

	return nil
}
