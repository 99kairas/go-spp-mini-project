package models

import "github.com/google/uuid"

type Grade struct {
	ID    uuid.UUID `json:"id"`
	Level string    `json:"level" form:"level"`
}
