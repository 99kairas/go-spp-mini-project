package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	ID            uuid.UUID `json:"id" form:"id"`
	SppID         uuid.UUID `json:"spp_id" form:"spp_id" gorm:"foreignKey:SppID;size:191"`
	Spp           SPP
	StudentID     uuid.UUID `json:"student_id" form:"student_id" gorm:"foreignKey:StudentID;size:191"`
	Student       Student
	AdminID       uuid.UUID `json:"admin_id" form:"admin_id" gorm:"foreignKey:AdminID;size:191"`
	Admin         Admin
	TotalAmount   float64    `json:"total_amount" form:"total_amount"`
	PaymentDate   *time.Time `json:"payment_date" form:"payment_date"`
	PaymentPhoto  string     `json:"payment_photo" form:"payment_photo"`
	PaymentStatus bool       `json:"payment_status" form:"payment_status"`
}
