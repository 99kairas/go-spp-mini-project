package controllers

import (
	"go-spp/middlewares"
	"go-spp/models/payloads"
	"go-spp/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateSPPController(c echo.Context) error {
	if _, err := middlewares.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, payloads.Response{
			Message: "route only for admin",
		})
	}

	payloadSPP := payloads.CreateSPPRequest{}
	c.Bind(&payloadSPP)

	response, err := usecase.CreateSPP(&payloadSPP)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "error create spp",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "success create spp",
		Data:    response,
	})
}
