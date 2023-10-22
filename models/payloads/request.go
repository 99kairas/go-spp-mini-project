package payloads

import "github.com/google/uuid"

type CreateAdminRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Name     string `json:"name" form:"name"`
	Address  string `json:"address" form:"address"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

type LoginAdminRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required, min=6"`
}

type CreateStudentRequest struct {
	NIS         string    `json:"nis" form:"nis" validate:"required"`
	Password    string    `json:"password" form:"password" validate:"required,min=6"`
	Firstname   string    `json:"first_name" form:"first_name"`
	Lastname    string    `json:"last_name" form:"last_name"`
	PhoneNumber string    `json:"phone_number" form:"phone_number"`
	Address     string    `json:"address" form:"address"`
	GradeID     uuid.UUID `json:"grade_id" form:"grade_id"`
}

type LoginStudentRequest struct {
	NIS      string `json:"nis" form:"nis" validate:"required"`
	Password string `json:"password" form:"password" validate:"required, min=6"`
}

type CreateSPPRequest struct {
	GradeID uuid.UUID `json:"grade_id" form:"grade_id"`
	Year    string    `json:"year" form:"year"`
	Month   string    `json:"month" form:"month"`
	Amount  float64   `json:"amount" form:"amount"`
}
