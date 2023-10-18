package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	ID       uuid.UUID `json:"id" form:"id"`
	Username string    `json:"username" form:"username" gorm:"unique"`
	Password string    `json:"password" form:"password"`
	Name     string    `json:"name" form:"name"`
	Address  string    `json:"address" form:"address"`
	Token    string    `json:"-" form:"-" gorm:"-"`
}
