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

type CreateStudentResponse struct {
	ID        uuid.UUID `json:"id"`
	NIS       string    `json:"nomor_induk_siswa"`
	Firstname string    `json:"first_name"`
	Lastname  string    `json:"last_name"`
	Address   string    `json:"address"`
	Class     string    `json:"class"`
}

type LoginStudentResponse struct {
	NIS   string `json:"nis"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
