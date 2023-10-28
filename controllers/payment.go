package controllers

import (
	"go-spp/middlewares"
	"go-spp/models/payloads"
	"go-spp/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminCreatePaymentByStudentIDController(c echo.Context) error {
	if _, err := middlewares.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, payloads.Response{
			Message: "route only for admin",
		})
	}

	payloadPayment := payloads.AdminCreatePaymentByStudentIDRequest{}
	c.Bind(&payloadPayment)

	response, err := usecase.CreatePaymentByStudentID(&payloadPayment)
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

func AdminCreatePaymentAllStudentController(c echo.Context) error {
	if _, err := middlewares.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, payloads.Response{
			Message: "route only for admin",
		})
	}

	payloadPayment := payloads.AdminCreatePaymentAllStudentRequest{}
	c.Bind(&payloadPayment)

	response, err := usecase.CreatePaymentAllStudent(&payloadPayment)
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

func GetAllPaymentsController(c echo.Context) error {
	if _, err := middlewares.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, payloads.Response{
			Message: "route only for admin",
		})
	}

	payment, err := usecase.GetAllPayments()
	if err != nil {
		return c.JSON(http.StatusNotFound, payloads.Response{
			Message: "payment not found",
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "success get all data",
		Data:    payment,
	})
}
