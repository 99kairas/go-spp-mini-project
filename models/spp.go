package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SPP struct {
	gorm.Model
	ID      uuid.UUID `json:"id" form:"id"`
	Grade   Grade
	GradeID uuid.UUID `json:"grade_id" form:"grade_id" gorm:"foreignKey:GradeID;size:191"`
	Year    string    `json:"year" form:"year"`
	Month   string    `json:"month" form:"month"`
	Amount  float64   `json:"amount" form:"amount"`
}
