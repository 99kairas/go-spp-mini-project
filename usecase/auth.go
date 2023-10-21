package usecase

import (
	"errors"
	"go-spp/middlewares"
	"go-spp/models/payloads"
	"go-spp/repositories"

	"golang.org/x/crypto/bcrypt"
)

func LoginAdmin(req *payloads.LoginAdminRequest) (res payloads.LoginAdminResponse, err error) {
	admin, err := repositories.GetAdmin(req.Username)
	if err != nil {
		return res, errors.New("username is not registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return res, errors.New("wrong password")
	}

	token, err := middlewares.CreateTokenAdmin(admin.ID)
	if err != nil {
		return res, errors.New("failed to create token")
	}

	admin.Token = token

	res = payloads.LoginAdminResponse{
		Username: admin.Username,
		Token:    admin.Token,
	}

	return
}

func LoginStudent(req *payloads.LoginStudentRequest) (res payloads.LoginStudentResponse, err error) {
	student, err := repositories.GetStudent(req.NIS)
	if err != nil {
		return res, errors.New("username is not registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(req.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return res, errors.New("wrong password")
	}

	token, err := middlewares.CreateTokenUser(student.ID)
	if err != nil {
		return res, errors.New("failed to create token")
	}

	student.Token = token

	res = payloads.LoginStudentResponse{
		NIS:   student.NIS,
		Name:  student.FirstName + " " + student.LastName,
		Token: student.Token,
	}

	return
}
