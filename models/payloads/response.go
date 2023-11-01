package payloads

import (
	"go-spp/models"

	"time"

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

type GetStudentByID struct {
	ID             uuid.UUID    `json:"id"`
	NIS            string       `json:"nis"`
	Name           string       `json:"name"`
	BirthDate      *time.Time   `json:"birth_date"`
	PhoneNumber    string       `json:"phone_number"`
	Address        string       `json:"address"`
	ProfilePicture string       `json:"profile_picture"`
	Grade          models.Grade `json:"grade"`
	CreatedAt      *time.Time   `json:"created_at"`
	UpdatedAt      *time.Time   `json:"updated_at"`
}

type GetAllStudent struct {
	ID             uuid.UUID    `json:"id"`
	NIS            string       `json:"nis"`
	Name           string       `json:"name"`
	BirthDate      *time.Time   `json:"birth_date"`
	PhoneNumber    string       `json:"phone_number"`
	Address        string       `json:"address"`
	ProfilePicture string       `json:"profile_picture"`
	Grade          models.Grade `json:"grade"`
	CreatedAt      *time.Time   `json:"created_at"`
	UpdatedAt      *time.Time   `json:"updated_at"`
}

type AdminCreatePaymentByStudentIDResponse struct {
	ID            uuid.UUID `json:"id"`
	SppID         uuid.UUID `json:"spp_id"`
	StudentID     uuid.UUID `json:"student_id"`
	AdminID       uuid.UUID `json:"admin_id"`
	TotalAmount   float64   `json:"total_amount"`
	PaymentStatus bool      `json:"payment_status"`
}

type GetAllPaymentsResponse struct {
	ID            uuid.UUID       `json:"id"`
	Spp           SPPResponse     `json:"spp"`
	Student       StudentResponse `json:"student"`
	Admin         AdminResponse   `json:"admin"`
	TotalAmount   float64         `json:"total_amount"`
	PaymentPhoto  string          `json:"payment_photo"`
	PaymentStatus bool            `json:"payment_status"`
	CreatedAt     *time.Time      `json:"created_at"`
	UpdatedAt     *time.Time      `json:"updated_at"`
}

type StudentResponse struct {
	ID   uuid.UUID `json:"id"`
	NIS  string    `json:"nis"`
	Name string    `json:"name"`
}

type AdminResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
}

type SPPResponse struct {
	ID    uuid.UUID `json:"id"`
	Year  string    `json:"year"`
	Month string    `json:"month"`
}

type AdminCreatePaymentAllStudentResponse struct {
	ID            uuid.UUID `json:"id"`
	SppID         uuid.UUID `json:"spp_id"`
	GradeID       uuid.UUID `json:"grade_id"`
	AdminID       uuid.UUID `json:"admin_id"`
	TotalAmount   float64   `json:"total_amount"`
	PaymentStatus bool      `json:"payment_status"`
}

type UpdateProfileStudentResponse struct {
	BirthDate   *time.Time `json:"birth_date"`
	PhoneNumber string     `json:"phone_number"`
	Address     string     `json:"address"`
}

type ApproveRejectPaymentResponse struct {
	PaymentID     uuid.UUID `json:"payment_id"`
	AdminID       uuid.UUID `json:"admin_id"`
	PaymentStatus bool      `json:"payment_status"`
}
