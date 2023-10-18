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
	Class          string     `json:"class" form:"class"`
	Address        string     `json:"address" form:"address"`
	ProfilePicture string     `json:"profile_picture" form:"profile_picture"`
	Token          string     `json:"-" form:"-" gorm:"-"`
}
