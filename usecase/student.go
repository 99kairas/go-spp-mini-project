package usecase

import (
	"errors"
	"go-spp/models"
	"go-spp/models/payloads"
	"go-spp/repositories"

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
		Class:       req.Class,
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
		Class:     newStudent.Class,
	}

	return

}
