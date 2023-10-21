package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SPP struct {
	gorm.Model
	ID     uuid.UUID `json:"id" form:"id"`
	Year   string    `json:"year" form:"year"`
	Month  string    `json:"month" form:"month"`
	Amount float64   `json:"amount" form:"amount"`
}
