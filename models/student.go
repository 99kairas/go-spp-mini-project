package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ID             uuid.UUID  `json:"id" form:"id"`
	NIS            string     `json:"nis" form:"nis" gorm:"unique"`
	Password       string     `json:"password" form:"password"`
	FirstName      string     `json:"first_name" form:"first_name"`
	LastName       string     `json:"last_name" form:"last_name"`
	BirthDate      *time.Time `json:"birth_date" form:"birth_date"`
	PhoneNumber    string     `json:"phone_number" form:"phone_number"`
	Address        string     `json:"address" form:"address"`
	ProfilePicture string     `json:"profile_picture" form:"profile_picture"`
	Grade          Grade
	GradeID        uuid.UUID `json:"grade_id" form:"grade_id" gorm:"foreignKey:GradeID;size:191"`
	Token          string    `json:"-" form:"-" gorm:"-"`
}
