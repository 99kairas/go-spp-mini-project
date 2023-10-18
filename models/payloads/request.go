package payloads

type CreateAdminRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Name     string `json:"name" form:"name"`
	Address  string `json:"address" form:"address"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

type LoginAdminRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required, min=6"`
}
