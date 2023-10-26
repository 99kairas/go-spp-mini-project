package controllers

import (
	"go-spp/middlewares"
	"go-spp/models/payloads"
	"go-spp/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminCreatePaymentController(c echo.Context) error {
	if _, err := middlewares.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, payloads.Response{
			Message: "route only for admin",
		})
	}

	payloadPayment := payloads.AdminCreatePaymentRequest{}
	c.Bind(&payloadPayment)

	response, err := usecase.CreatePayment(&payloadPayment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "error create payment",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "success create payment",
		Data:    response,
	})
}
