package controllers

import (
	"go-spp/models/payloads"
	"go-spp/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateSPPController(c echo.Context) error {
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
