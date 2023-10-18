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
