package controllers

import (
	"go-spp/middlewares"
	"go-spp/models/payloads"
	"go-spp/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterStudentController(c echo.Context) error {
	if _, err := middlewares.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, payloads.Response{
			Message: "route only for admin",
		})
	}

	payloadUser := payloads.CreateStudentRequest{}
	c.Bind(&payloadUser)

	response, err := usecase.CreateStudent(&payloadUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "error create student",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "success register student",
		Data:    response,
	})
}

func LoginStudentController(c echo.Context) error {
	payloadUser := payloads.LoginStudentRequest{}

	c.Bind(&payloadUser)

	response, err := usecase.LoginStudent(&payloadUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "error login student",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "success login",
		Data:    response,
	})
}
