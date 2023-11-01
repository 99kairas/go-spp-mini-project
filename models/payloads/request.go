package payloads

import (
	"time"

	"github.com/google/uuid"
)

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

type AdminCreatePaymentByStudentIDRequest struct {
	SppID         uuid.UUID  `json:"spp_id" form:"spp_id"`
	StudentID     uuid.UUID  `json:"student_id" form:"student_id"`
	AdminID       uuid.UUID  `json:"admin_id" form:"admin_id"`
	TotalAmount   float64    `json:"total_amount" form:"total_amount"`
	PaymentDate   *time.Time `json:"payment_date" form:"payment_date"`
	PaymentPhoto  string     `json:"payment_photo" form:"payment_photo"`
	PaymentStatus bool       `json:"payment_status" form:"payment_status"`
}

type AdminCreatePaymentAllStudentRequest struct {
	SppID         uuid.UUID  `json:"spp_id" form:"spp_id"`
	GradeID       uuid.UUID  `json:"grade_id" form:"grade_id"`
	AdminID       uuid.UUID  `json:"admin_id" form:"admin_id"`
	TotalAmount   float64    `json:"total_amount" form:"total_amount"`
	PaymentDate   *time.Time `json:"payment_date" form:"payment_date"`
	PaymentPhoto  string     `json:"payment_photo" form:"payment_photo"`
	PaymentStatus bool       `json:"payment_status" form:"payment_status"`
}

type UpdatePasswordRequest struct {
	CurrentPassword string `json:"current_password" form:"current_password"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type UpdateProfileStudentRequest struct {
	BirthDate   string `json:"birth_date" form:"birth_date"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Address     string `json:"address" form:"address"`
}

type UploadPaymentPhotoRequest struct {
	PaymentID uuid.UUID `json:"payment_id" form:"payment_id"`
	Image     string    `json:"image" form:"image"`
}

type ApproveRejectPaymentRequest struct {
	PaymentID uuid.UUID `json:"payment_id" form:"payment_id"`
	AdminID   uuid.UUID `json:"admin_id" form:"admin_id"`
}
