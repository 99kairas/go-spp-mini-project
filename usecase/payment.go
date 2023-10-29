package usecase

import (
	"errors"
	"go-spp/models"
	"go-spp/models/payloads"
	"go-spp/repositories"

	"github.com/google/uuid"
)

func GetAllPayments() (res []payloads.GetAllPaymentsResponse, err error) {
	payments, err := repositories.GetAllPayments()
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

func GetPaymentsWithPhoto() (res []payloads.GetAllPaymentsResponse, err error) {
	payments, err := repositories.GetAllPayments()
	if err != nil {
		return nil, errors.New("payments not found")
	}

	for _, student := range payments {
		if student.PaymentPhoto != "" {
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
	}

	return res, nil
}

func CreatePaymentByStudentID(req *payloads.AdminCreatePaymentByStudentIDRequest) (res payloads.AdminCreatePaymentByStudentIDResponse, err error) {
	if !repositories.IsPaymentAvailableByStudentID(req.SppID) {
		return res, errors.New("payment is already created")
	}

	spp, err := repositories.GetSPPByID(req.SppID)
	if err != nil {
		return res, err
	}

	newPayment := &models.Payment{
		ID:            uuid.New(),
		SppID:         req.SppID,
		StudentID:     req.StudentID,
		AdminID:       req.AdminID,
		TotalAmount:   spp.Amount,
		PaymentPhoto:  "",
		PaymentStatus: false,
	}

	err = repositories.CreatePaymentByStudentID(newPayment)

	if err != nil {
		return
	}

	res = payloads.AdminCreatePaymentByStudentIDResponse{
		ID:            newPayment.ID,
		SppID:         newPayment.SppID,
		StudentID:     newPayment.StudentID,
		AdminID:       newPayment.AdminID,
		TotalAmount:   newPayment.TotalAmount,
		PaymentStatus: newPayment.PaymentStatus,
	}

	return
}

func CreatePaymentAllStudent(req *payloads.AdminCreatePaymentAllStudentRequest) (res payloads.AdminCreatePaymentAllStudentResponse, err error) {
	students, err := repositories.IsStudentAvailable(req.GradeID)
	if err != nil {
		return res, err
	}

	spp, err := repositories.GetSPPByID(req.SppID)
	if err != nil {
		return res, err
	}

	var newPayment *models.Payment

	for _, student := range students {
		paymentExists, err := repositories.IsPaymentAvailable(student.ID, req.SppID)
		if err != nil {
			return res, err
		}

		if !paymentExists {
			newPayment = &models.Payment{
				ID:            uuid.New(),
				SppID:         req.SppID,
				StudentID:     student.ID,
				AdminID:       req.AdminID,
				TotalAmount:   spp.Amount,
				PaymentPhoto:  "",
				PaymentStatus: false,
			}

			err := repositories.CreatePaymentAllStudent(newPayment)
			if err != nil {
				return res, err
			}
		} else {
			return res, errors.New("payment is already created")
		}
	}

	if newPayment != nil {
		res = payloads.AdminCreatePaymentAllStudentResponse{
			ID:            newPayment.ID,
			SppID:         newPayment.SppID,
			GradeID:       req.GradeID,
			AdminID:       req.AdminID,
			TotalAmount:   newPayment.TotalAmount,
			PaymentStatus: newPayment.PaymentStatus,
		}
	}

	return res, nil
}
