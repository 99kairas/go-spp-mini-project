package controllers

import (
	"go-spp/models/payloads"
	"go-spp/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterAdminController(c echo.Context) error {
	payloadUser := payloads.CreateAdminRequest{}
	c.Bind(&payloadUser)

	response, err := usecase.CreateAdmin(&payloadUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "error create admin",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "success register admin",
		Data:    response,
	})
}

func LoginAdminController(c echo.Context) error {
	payloadUser := payloads.LoginAdminRequest{}

	c.Bind(&payloadUser)

	response, err := usecase.LoginAdmin(&payloadUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "error login admin",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "success login as admin",
		Data:    response,
	})
}
