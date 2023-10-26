package payloads

import (
	"go-spp/models"

	"github.com/google/uuid"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CreateAdminResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
	Address  string    `json:"address"`
}

type LoginAdminResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type CreateStudentResponse struct {
	ID        uuid.UUID `json:"id"`
	NIS       string    `json:"nomor_induk_siswa"`
	Firstname string    `json:"first_name"`
	Lastname  string    `json:"last_name"`
	Address   string    `json:"address"`
	GradeID   uuid.UUID `json:"grade_id"`
}

type LoginStudentResponse struct {
	NIS   string `json:"nis"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type CreateSPPResponse struct {
	ID      uuid.UUID `json:"id"`
	GradeID uuid.UUID `json:"grade_id"`
	Year    string    `json:"year"`
	Month   string    `json:"month"`
	Amount  float64   `json:"amount"`
}

type AdminCreatePaymentResponse struct {
	ID            uuid.UUID `json:"id"`
	SppID         uuid.UUID `json:"spp_id"`
	Spp           []models.SPP
	StudentID     uuid.UUID `json:"student_id"`
	AdminID       uuid.UUID `json:"admin_id"`
	TotalAmount   float64   `json:"total_amount"`
	PaymentStatus bool      `json:"payment_status"`
}

type GetAllPaymentsResponse struct {
	ID            uuid.UUID `json:"id"`
	Spp           []models.SPP
	Student       []models.Student
	Admin         []models.Admin
	TotalAmount   float64 `json:"total_amount"`
	PaymentPhoto  string  `json:"payment_photo"`
	PaymentStatus bool    `json:"payment_status"`
}
