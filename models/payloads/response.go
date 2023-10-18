package payloads

import "github.com/google/uuid"

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
