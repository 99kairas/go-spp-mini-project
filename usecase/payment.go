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

func GetPaymentByID(paymentID uuid.UUID) (res []payloads.GetAllPaymentsResponse, err error) {
	payments, err := repositories.GetPaymentByID(paymentID)
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
				GradeID:       req.GradeID,
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
			GradeID:       newPayment.GradeID,
			AdminID:       newPayment.AdminID,
			TotalAmount:   newPayment.TotalAmount,
			PaymentStatus: newPayment.PaymentStatus,
		}
	} else {
		return res, errors.New("bad request")
	}

	return res, nil
}

func ApprovePayment(paymentID uuid.UUID, adminID uuid.UUID) (res payloads.ApproveRejectPaymentResponse, err error) {
	payment, err := repositories.GetPaymentID(paymentID)
	if err != nil {
		return res, errors.New("payment not found")
	}

	payment.PaymentStatus = true

	payment.AdminID = adminID

	err = repositories.UpdatePayment(payment)
	if err != nil {
		return res, errors.New("can't update payment")
	}

	res = payloads.ApproveRejectPaymentResponse{
		PaymentID:     paymentID,
		AdminID:       adminID,
		PaymentStatus: payment.PaymentStatus,
	}

	return res, nil
}

func RejectPayment(paymentID uuid.UUID, adminID uuid.UUID) (res payloads.ApproveRejectPaymentResponse, err error) {
	payment, err := repositories.GetPaymentID(paymentID)
	if err != nil {
		return res, errors.New("payment not found")
	}

	payment.PaymentStatus = false

	payment.AdminID = adminID

	payment.PaymentPhoto = " "

	err = repositories.UpdatePayment(payment)
	if err != nil {
		return res, errors.New("can't update payment")
	}

	res = payloads.ApproveRejectPaymentResponse{
		PaymentID:     paymentID,
		AdminID:       adminID,
		PaymentStatus: payment.PaymentStatus,
	}

	return res, nil
}
