package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	ID            uuid.UUID  `json:"id" form:"id"`
	SppID         uuid.UUID  `json:"spp_id" form:"spp_id" gorm:"foreignKey:SppID;type:char(40);unique"`
	Spp           SPP        `json:"spp" form:"spp"`
	NIS           uuid.UUID  `json:"nis_id" form:"nis_id" gorm:"foreignKey:SppID;type:char(40);unique"`
	Student       Student    `json:"student" form:"student"`
	AdminID       uuid.UUID  `json:"admin_id" form:"admin_id" gorm:"foreignKey:SppID;type:char(40);unique"`
	Admin         Admin      `json:"admin" form:"admin"`
	TotalAmount   float64    `json:"total_amount" form:"total_amount"`
	PaymentDate   *time.Time `json:"payment_date" form:"payment_date"`
	PaymentPhoto  string     `json:"payment_photo" form:"payment_photo"`
	PaymentStatus bool       `json:"payment_status" form:"payment_status"`
}
