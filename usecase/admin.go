package usecase

import (
	"errors"
	"go-spp/models"
	"go-spp/models/payloads"
	"go-spp/repositories"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateAdmin(req *payloads.CreateAdminRequest) (resp payloads.CreateAdminResponse, err error) {
	if !repositories.IsUsernameAvailable(req.Username) {
		return resp, errors.New("username is already registered")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return
	}

	newAdmin := &models.Admin{
		ID:       uuid.New(),
		Username: req.Username,
		Password: string(passwordHash),
		Name:     req.Name,
		Address:  req.Address,
	}

	err = repositories.CreateAdmin(newAdmin)

	if err != nil {
		return
	}

	resp = payloads.CreateAdminResponse{
		ID:       newAdmin.ID,
		Username: newAdmin.Username,
		Name:     newAdmin.Name,
		Address:  newAdmin.Address,
	}

	return

}

func CreateSPP(req *payloads.CreateSPPRequest) (res payloads.CreateSPPResponse, err error) {
	if !repositories.IsSPPAvailable(req.Year, req.Month) {
		return res, errors.New("spp is already created")
	}
	newSPP := &models.SPP{
		ID:      uuid.New(),
		GradeID: req.GradeID,
		Year:    req.Year,
		Month:   req.Month,
		Amount:  req.Amount,
	}

	err = repositories.CreateSPP(newSPP)

	if err != nil {
		return
	}

	res = payloads.CreateSPPResponse{
		ID:      newSPP.ID,
		GradeID: newSPP.GradeID,
		Year:    newSPP.Year,
		Month:   newSPP.Month,
		Amount:  newSPP.Amount,
	}

	return
}

func CreatePayment(req *payloads.AdminCreatePaymentRequest) (res payloads.AdminCreatePaymentResponse, err error) {
	if !repositories.IsPaymentAvailable(req.SppID) {
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

	err = repositories.CreatePayment(newPayment)

	if err != nil {
		return
	}

	res = payloads.AdminCreatePaymentResponse{
		ID:            newPayment.ID,
		SppID:         newPayment.SppID,
		StudentID:     newPayment.StudentID,
		AdminID:       newPayment.AdminID,
		TotalAmount:   newPayment.TotalAmount,
		PaymentStatus: newPayment.PaymentStatus,
	}

	return
}
